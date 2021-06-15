package util

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/xiyouhpy/spider/base"

	"gopkg.in/yaml.v2"
)

const (
	// mysql 相关配置文件
	MysqlConf = "./config/mysql.yaml"
	// redis 相关配置文件
	RedisConf = "./config/redis.yaml"

	// 业务相关配置文件
	SpiderConf = "./config/spider.yaml"
)

// mysql配置信息
type mysqlInfo struct {
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	ConnTimeout  int    `yaml:"conn_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
	ReadTimeout  int    `yaml:"read_timeout"`
	Retry        int    `yaml:"retry"`
}

var MysqlConfInfo = map[string]mysqlInfo{}

// redis配置信息
type redisInfo struct {
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	ConnTimeout  int    `yaml:"conn_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
	ReadTimeout  int    `yaml:"read_timeout"`
	Retry        int    `yaml:"retry"`
}

var RedisConfInfo = map[string]redisInfo{}

func InitConf() {
	// 初始化mysql配置信息
	if getMysqlConf(MysqlConf, &MysqlConfInfo) != base.ErrSuccess {
		return
	}

	// 初始化redis配置信息
	if getRedisConf(RedisConf, &RedisConfInfo) != base.ErrSuccess {
		return
	}
}

// getMysqlConf 获取配置信息
func getMysqlConf(strFileName string, confInfo *map[string]mysqlInfo) error {
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

// getRedisConf 获取配置信息
func getRedisConf(strFileName string, confInfo *map[string]redisInfo) error {
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
