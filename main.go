package main

import (
	"context"

	"github.com/sirupsen/logrus"
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

var ctx = context.Background()

func main() {
	logrus.Info("test")
	router.Server()
	return
}
