package main

import (
	"github.com/xiyouhpy/spider/base"
	"github.com/xiyouhpy/spider/router"
)

func init() {
	// 获取日志信息
	base.Logger()
}

func main() {
	router.Server()
	return
}
