package base

import (
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var LogPath = "./log/spider.log"

func logWriter(path string, level string) *rotatelogs.RotateLogs {
	fileSuffix := ".info"
	if level == "error" || level == "fatal" || level == "panic" {
		fileSuffix = ".error"
	}

	logier, err := rotatelogs.New(
		path+fileSuffix+"."+"%Y%m%d%H%M",
		rotatelogs.WithLinkName(path),              // 生成软链，指向最新日志文件
		rotatelogs.WithRotationCount(240),          // 文件最大保存份数
		rotatelogs.WithRotationTime(time.Minute*1), // 日志切割时间间隔
	)

	if err != nil {
		panic(err)
	}
	return logier
}

func newLogger(log *logrus.Logger, logPath string) {

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: logWriter(logPath, "debug"), // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  logWriter(logPath, "info"),
		logrus.WarnLevel:  logWriter(logPath, "warn"),
		logrus.ErrorLevel: logWriter(logPath, "error"),
		logrus.FatalLevel: logWriter(logPath, "fatal"),
		logrus.PanicLevel: logWriter(logPath, "panic"),
	}, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	log.AddHook(lfHook)
}

func InitLog(log *logrus.Logger) {
	log.Out = os.Stdout
	var loglevel logrus.Level
	err := loglevel.UnmarshalText([]byte("info"))
	if err != nil {
		log.Panicf("设置log级别失败：%v", err)
	}

	log.SetLevel(loglevel)

	newLogger(log, LogPath)
}
