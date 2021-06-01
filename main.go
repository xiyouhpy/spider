package main

import (
	"github.com/xiyouhpy/spider/util"

	"log"
)

func main() {
	util.InitConf()
	log.Printf("spider info:%+v", util.SpiderConfInfo)
	log.Printf("system info:%+v", util.SystemConfInfo)
	log.Printf("mysql info:%+v", util.MysqlConfInfo)
	log.Printf("redis info:%+v", util.RedisConfInfo)
	return
}
