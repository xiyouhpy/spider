package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xiyouhpy/spider/base"
)

// GetListHandler ...
func GetList(c *gin.Context) {
	message := c.Query("message")

	data := map[string]string{
		"message": message,
	}
	JsonRet(c, base.ErrSuccess, data)
}
