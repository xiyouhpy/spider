package main

import (
	"log"

	"github.com/xiyouhpy/spider/router"
	"github.com/xiyouhpy/spider/util"
)

func main() {
	util.InitConf()
	log.Printf("spider info:%+v", util.SpiderConfInfo)
	log.Printf("system info:%+v", util.SystemConfInfo)
	log.Printf("mysql info:%+v", util.MysqlConfInfo)
	log.Printf("redis info:%+v", util.RedisConfInfo)

	router.Server()
	return
}
