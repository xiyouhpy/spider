package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xiyouhpy/spider/base"
	"github.com/xiyouhpy/spider/service"
	"github.com/xiyouhpy/spider/util"
)

// GetList 获取参数列表信息
func GetList(c *gin.Context) {
	var req util.GetListReq
	req.OPtime = c.DefaultQuery("op_time", "")
	req.Ctime = c.DefaultQuery("ctime", "")
	req.Size = c.DefaultQuery("size", "")
	req.Offset = c.DefaultQuery("pn", "0")
	req.Limit = c.DefaultQuery("rn", "20")

	res, err := service.GetList(ctx, &req)
	if err != nil {
		JsonRet(c, base.ErrServiceError, nil)
		return
	}

	JsonRet(c, base.ErrSuccess, res)
}
