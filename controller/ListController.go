package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xiyouhpy/spider/base"
	"github.com/xiyouhpy/spider/service"
)

// GetListHandler ...
func GetList(c *gin.Context) {
	serviceID := c.Query("message")

	req := map[string]interface{}{
		"service_id": serviceID,
	}
	res, err := service.GetList(ctx, req)
	if err == nil {
		JsonRet(c, base.ErrServiceError, nil)
		return
	}

	ret := map[string]interface{}{
		"service_id": res,
	}
	JsonRet(c, base.ErrSuccess, ret)
}
