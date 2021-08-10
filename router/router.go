package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/xiyouhpy/spider/controller"
	"github.com/xiyouhpy/spider/util"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// tmpFileExpireTime 临时文件过期时间设置
const tmpFileExpireTime = 3600

// Server ...
func Server() {
	r := gin.Default()

	// 创建必要的目录
	initPath()

	// 清理过期文件
	go cleanPath()

	// 路由注册和跳转
	registerService(r)

	// 服务监听端口
	r.Run(":8000")
}

// registerService ...
func registerService(r *gin.Engine) {
	// 抓取图片接口
	r.GET("/spider/grabImage", controller.GrabImage)
}

// initPath ...
func initPath() bool {
	for _, tmpDir := range util.ArrDirs {
		// 检查目录是否存在
		if _, err := os.Stat(tmpDir); !os.IsNotExist(err) {
			continue
		}

		// 创建目录
		err := os.MkdirAll(tmpDir, os.ModePerm)
		if err != nil {
			logrus.Warnf("initPath create err, dir:%s, err:%s", tmpDir, err.Error())
			return false
		}
	}

	return true
}

// cleanPath ...
func cleanPath() {
	ticker := time.NewTicker(time.Minute * time.Duration(10))
	for range ticker.C {
		intNow := time.Now().Unix() - tmpFileExpireTime
		for _, tmpDir := range util.ArrDirs {
			fileList, listErr := ioutil.ReadDir(tmpDir)
			if listErr != nil {
				logrus.Errorf("ioutil.ReadDir err, tmpDir:%s, err:%s", tmpDir, listErr.Error())
				continue
			}

			for _, file := range fileList {
				fileName := file.Name()
				if fileName != "." && fileName != ".." && file.ModTime().Unix() < intNow {
					err := os.RemoveAll(filepath.Join(tmpDir, "/", fileName))
					logrus.Infof("CleanFile RemoveAll, cleanName:%s, err:%+v", fileName, err)
				}
			}
		}
	}
}
