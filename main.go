package main

import (
	"log"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/xiyouhpy/spider/router"
	"github.com/xiyouhpy/spider/util"
)

func init() {
	// 设置日志保存目录
	logPath := "./log/spider.log"
	// 设置日志切割时间间隔，单位：小时
	rotationTime := 1
	// 设置日志保留个数
	rotationCount := 240

	writer, _ := rotatelogs.New(
		logPath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(logPath),
		rotatelogs.WithRotationCount(uint(rotationCount)),
		rotatelogs.WithRotationTime(time.Duration(rotationTime)*time.Hour),
	)
	logrus.SetOutput(writer)
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

func main() {
	util.InitConf()
	log.Printf("spider info:%+v", util.SpiderConfInfo)
	log.Printf("system info:%+v", util.SystemConfInfo)
	log.Printf("mysql info:%+v", util.MysqlConfInfo)
	log.Printf("redis info:%+v", util.RedisConfInfo)

	router.Server()
	return
}
