package main

import (
	"github.com/sirupsen/logrus"
	"github.com/xiyouhpy/spider/base"
	"github.com/xiyouhpy/spider/router"
	"github.com/xiyouhpy/spider/util"
	"log"
	"time"
)

func init() {
	// 初始化配置信息
	util.InitConf()

	// 初始化日志信息
	base.InitLog()
}

func main() {
	log.Printf("mysql info:%+v", util.MysqlConfInfo)
	log.Printf("redis info:%+v", util.RedisConfInfo)

	for {
		logrus.Infof("test %s")
		time.Sleep(2 * time.Second)
	}
	router.Server()
	return
}
