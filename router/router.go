package router

import (
	"github.com/xiyouhpy/spider/controller"

	"github.com/gin-gonic/gin"
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
