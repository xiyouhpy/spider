package util

import (
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/xiyouhpy/spider/base"
	"gopkg.in/yaml.v2"
)

const (
	// MysqlConf 相关配置文件
	MysqlConf = "./config/mysql.yaml"
	// RedisConf 相关配置文件
	RedisConf = "./config/redis.yaml"
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

// GetConf 获取配置信息
func GetConf() {
	// 获取 mysql 配置信息
	if getMysqlConf(MysqlConf, &MysqlConfInfo) != base.ErrSuccess {
		return
	}

	// 获取 redis 配置信息
	if getRedisConf(RedisConf, &RedisConfInfo) != base.ErrSuccess {
		return
	}
}

// getMysqlConf 获取配置信息
func getMysqlConf(strFileName string, confInfo *map[string]mysqlInfo) error {
	_, err := os.Stat(strFileName)
	if os.IsNotExist(err) {
		logrus.Fatalf("getSpiderConf err, filename:%s not exist", strFileName)
		return base.ErrGetConfError
	}

	data, err := ioutil.ReadFile(strFileName)
	if err != nil {
		logrus.Fatalf("getSpiderConf ioutil.ReadFile err, filename:%s, err:%s", strFileName, err.Error())
		return base.ErrGetConfError
	}

	err = yaml.Unmarshal(data, confInfo)
	if err != nil {
		logrus.Fatalf("getSpiderConf yaml.Unmarshal err, filename:%s, err:%s", strFileName, err.Error())
		return base.ErrGetConfError
	}

	return base.ErrSuccess
}

// getRedisConf 获取配置信息
func getRedisConf(strFileName string, confInfo *map[string]redisInfo) error {
	_, err := os.Stat(strFileName)
	if os.IsNotExist(err) {
		logrus.Fatalf("getSpiderConf err, filename:%s not exist", strFileName)
		return base.ErrGetConfError
	}

	data, err := ioutil.ReadFile(strFileName)
	if err != nil {
		logrus.Fatalf("getSpiderConf ioutil.ReadFile err, filename:%s, err:%s", strFileName, err.Error())
		return base.ErrGetConfError
	}

	err = yaml.Unmarshal(data, confInfo)
	if err != nil {
		logrus.Fatalf("getSpiderConf yaml.Unmarshal err, filename:%s, err:%s", strFileName, err.Error())
		return base.ErrGetConfError
	}

	return base.ErrSuccess
}
