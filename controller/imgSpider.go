package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xiyouhpy/spider/base"
	"github.com/xiyouhpy/spider/service/deliver"
	"github.com/xiyouhpy/spider/util"
	"path/filepath"
	"time"
)

// GrabImage 抓取图片接口
func GrabImage(c *gin.Context) {
	srcUrl := c.DefaultQuery("src_url", "")
	if srcUrl == "" {
		JsonRet(c, base.ErrParamsError)
		return
	}

	intTime := time.Now().Unix()
	strMd5 := base.GetMd5(srcUrl)
	strExt := filepath.Ext(srcUrl)
	if strExt == "" {
		strExt = ".jpg"
	}

	dstName := fmt.Sprintf("%sspider_%d_%s", util.SpiderImgDir, intTime, strMd5[len(strMd5)-10:]+strExt)
	err := deliver.Download(srcUrl, dstName)
	if err != nil {
		JsonRet(c, base.ErrDownloadError)
		return
	}

	JsonRet(c, base.ErrSuccess, dstName)
	return
}


