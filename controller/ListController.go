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
	req.Atime = c.DefaultQuery("atime", "")
	req.Ctime = c.DefaultQuery("ctime", "")
	req.Size = c.DefaultQuery("size", "")
	req.Offset = c.DefaultQuery("offset", "0")
	req.Limit = c.DefaultQuery("limit", "20")

	res, err := service.GetList(ctx, &req)
	if err != nil {
		JsonRet(c, base.ErrCallServiceError, nil)
	} else {
		JsonRet(c, base.ErrSuccess, res)
	}
}

// AddList 获取参数列表信息
func AddList(c *gin.Context) {
	var req util.AddListReq
	req.Title = c.DefaultQuery("title", "")
	req.Abstract = c.DefaultQuery("abstract", "")
	req.Url = c.DefaultQuery("url", "")
	req.Path = c.DefaultQuery("path", "")
	req.Ctime = c.DefaultQuery("ctime", "")
	req.Size = c.DefaultQuery("size", "")
	req.Md5 = c.DefaultQuery("md5", "")

	if req.Title == "" || req.Url == "" || req.Path == "" || req.Ctime == "" || req.Size == "" || req.Md5 == "" {
		JsonRet(c, base.ErrParamsError, nil)
		return
	}

	res, err := service.AddList(ctx, &req)
	if err != nil {
		JsonRet(c, base.ErrCallServiceError, nil)
		return
	} else {
		JsonRet(c, base.ErrSuccess, res)
		return
	}
}

// DelList 删除参数列表信息
func DelList(c *gin.Context) {
	strId := c.DefaultQuery("id", "0")
	if strId == "" {
		JsonRet(c, base.ErrParamsError, nil)
		return
	}

	err := service.DelList(ctx, strId, "")
	if err != nil {
		JsonRet(c, base.ErrCallServiceError, nil)
		return
	} else {
		JsonRet(c, base.ErrSuccess, nil)
		return
	}
}
