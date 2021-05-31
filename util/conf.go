package util

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	// 系统相关配置文件
	LogFileConf = "./config/log.yaml"
	// mysql 相关配置文件
	MysqlConf = "./config/mysql.yaml"
	// redis 相关配置文件
	RedisConf = "./config/redis.yaml"

	// 业务相关配置文件
	SpiderFileConf = "./config/spider.yaml"
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
	logErr := getConf(LogFileConf, &SystemConfInfo)
	if logErr != nil {
		log.Fatalf("InitConf getConf log err, err:%s", logErr.Error())
		return
	}

	// 初始化抓取配置信息
	spiderErr := getConf(SpiderFileConf, &SpiderConfInfo)
	if spiderErr != nil {
		log.Fatalf("InitConf getConf spider err, err:%s", spiderErr.Error())
		return
	}

	// 初始化mysql配置信息
	mysqlErr := getConf(MysqlConf, &MysqlConfInfo)
	if mysqlErr != nil {
		log.Fatalf("InitConf getConf mysql err, err:%s", mysqlErr.Error())
		return
	}

	// 初始化redis配置信息
	redisErr := getConf(RedisConf, &RedisConfInfo)
	if redisErr != nil {
		log.Fatalf("InitConf getConf redis err, err:%s", redisErr.Error())
		return
	}
}

// getConf 获取配置信息
func getConf(strFileName string, confInfo *map[string]interface{}) error {
	_, err := os.Stat(strFileName)
	if os.IsNotExist(err) {
		log.Fatal("getSpiderConf err, file is not exist, filename：", strFileName)
		return nil
	}
	data, err := ioutil.ReadFile(strFileName)
	if err != nil {
		log.Fatal("getSpiderConf ioutil.ReadFile err, filename:", strFileName, " err:", err.Error())
		return err
	}

	err = yaml.Unmarshal(data, confInfo)
	if err != nil {
		log.Fatal("getSpiderConf yaml.Unmarshal err, filename:", strFileName, " err:", err.Error())
		return err
	}

	return nil
}
