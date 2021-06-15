package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xiyouhpy/spider/base"
	"github.com/xiyouhpy/spider/service"
)

// GetListHandler ...
func GetList(c *gin.Context) {
	message := c.Query("message")

	service.GetList(ctx, 120)
	data := map[string]string{
		"message": message,
	}
	JsonRet(c, base.ErrSuccess, data)
}
