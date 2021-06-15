package main

import (
	"context"

	"github.com/xiyouhpy/spider/base"
	"github.com/xiyouhpy/spider/router"
	"github.com/xiyouhpy/spider/util"
)

func init() {
	// 初始化配置信息
	util.InitConf()

	// 初始化日志信息
	base.InitLog()
}

var ctx = context.Background()

func main() {
	router.Server()
	return
}
