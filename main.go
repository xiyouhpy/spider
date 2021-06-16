package main

import (
	"github.com/xiyouhpy/spider/base"
	"github.com/xiyouhpy/spider/router"
	"github.com/xiyouhpy/spider/util"
)

func init() {
	// 获取配置信息
	util.GetConf()

	// 获取日志信息
	base.Logger()
}

func main() {
	router.Server()
	return
}
