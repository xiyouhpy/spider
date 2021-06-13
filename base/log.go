package base

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"time"
)

// 设置日志保存目录
const LogPath = "./log/spider.log"
// 设置日志切割时间间隔，单位：小时（每隔1小时切割）
const RotationTime = 1
// 设置日志保留个数（保留10天）
const RotationCount = 240

func InitLog() {
	writer, _ := rotatelogs.New(
		LogPath+".%Y%m%d%H",
		rotatelogs.WithLinkName(LogPath),
		rotatelogs.WithRotationCount(uint(RotationCount)),
		rotatelogs.WithRotationTime(time.Duration(RotationTime)*time.Hour),
	)
	logrus.SetOutput(writer)
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
}