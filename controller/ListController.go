package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xiyouhpy/spider/base"
	"github.com/xiyouhpy/spider/service"
)

// GetList 获取参数列表信息
func GetList(c *gin.Context) {
	serviceID := c.Query("service_id")

	req := map[string]interface{}{
		"service_id": serviceID,
	}
	res, err := service.GetList(ctx, req)
	if err != nil {
		JsonRet(c, base.ErrServiceError, nil)
		return
	}

	JsonRet(c, base.ErrSuccess, res)
}
