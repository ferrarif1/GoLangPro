package logger

import (
	"time"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func InitLog() *logrus.Logger {
	// 创建一个新的日志记录器
	log := logrus.New()

	// 设置日志记录级别
	log.SetLevel(logrus.InfoLevel)

	// 获取当前时间并将其格式化为字符串（包含毫秒）
	timestamp := time.Now().Format("2006-01-02-15-04-05.000")

	// 创建一个新的文件日志钩子
	logPath := "logs/" + timestamp + ".log"
	fileHook := lfshook.NewHook(lfshook.PathMap{
		logrus.InfoLevel:  logPath,
		logrus.ErrorLevel: logPath,
		logrus.WarnLevel:  logPath,
		logrus.DebugLevel: logPath,
		logrus.FatalLevel: logPath,
		logrus.PanicLevel: logPath,
	}, &logrus.JSONFormatter{})

	// 将日志钩子添加到记录器中
	log.AddHook(fileHook)

	return log
}
