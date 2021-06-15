package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xiyouhpy/spider/controller"
)

// Server ...
func Server() {
	r := gin.Default()

	// 路由注册和跳转
	registerService(r)

	// 服务监听端口
	r.Run(":8000")
}

// Router ...
func registerService(r *gin.Engine) {
	// 获取数据列表
	r.GET("/spider/getList", controller.GetList)
}
