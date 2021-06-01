package util

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/xiyouhpy/spider/base"

	"gopkg.in/yaml.v2"
)

const (
	// 系统相关配置文件
	LogConf = "./config/log.yaml"
	// mysql 相关配置文件
	MysqlConf = "./config/mysql.yaml"
	// redis 相关配置文件
	RedisConf = "./config/redis.yaml"

	// 业务相关配置文件
	SpiderConf = "./config/spider.yaml"
)

// 抓取配置信息
var SpiderConfInfo = map[string]interface{}{}

// 系统配置信息
var SystemConfInfo = map[string]interface{}{}

// mysql配置信息
var MysqlConfInfo = map[string]interface{}{}

// redis配置信息
var RedisConfInfo = map[string]interface{}{}

func InitConf() {
	// 初始化系统配置信息
	if getConf(LogConf, &SystemConfInfo) != base.ErrSuccess {
		return
	}

	// 初始化抓取配置信息
	if getConf(SpiderConf, &SpiderConfInfo) != base.ErrSuccess {
		return
	}

	// 初始化mysql配置信息
	if getConf(MysqlConf, &MysqlConfInfo) != base.ErrSuccess {
		return
	}

	// 初始化redis配置信息
	if getConf(RedisConf, &RedisConfInfo) != base.ErrSuccess {
		return
	}
}

// getConf 获取配置信息
func getConf(strFileName string, confInfo *map[string]interface{}) base.ErrCode {
	_, err := os.Stat(strFileName)
	if os.IsNotExist(err) {
		log.Fatal("getSpiderConf err, filename:", strFileName, " not exist")
		return base.ErrGetConfError
	}

	data, err := ioutil.ReadFile(strFileName)
	if err != nil {
		log.Fatal("getSpiderConf ioutil.ReadFile err, filename:", strFileName, " err:", err.Error())
		return base.ErrGetConfError
	}

	err = yaml.Unmarshal(data, confInfo)
	if err != nil {
		log.Fatal("getSpiderConf yaml.Unmarshal err, filename:", strFileName, " err:", err.Error())
		return base.ErrGetConfError
	}

	return base.ErrSuccess
}
