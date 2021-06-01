package router

import (
	"github.com/xiyouhpy/spider/controller"

	"github.com/gin-gonic/gin"
)

// MapRouter ...
func MapRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.GET("/spider/getList", controller.GetList)
	return router
}
