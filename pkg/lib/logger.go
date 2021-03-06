package lib

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type Logger struct {
	*logrus.Logger
}

// NewLogger 构造日志服务
func NewLogger(logPath, module string) (*Logger, error) {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return nil, err
	}
	// 定义文件前缀和日志名称
	prefix := logPath + "/" + module
	latestLogFile := prefix + ".log"

	logClient := logrus.New()
	logClient.Out = src
	// logClient.Out = os.Stdout // 是否打印到控制台上
	logClient.SetLevel(logrus.DebugLevel)

	logWriter, err := rotatelogs.New(
		prefix+"-%Y-%m-%d.log",                    // 生成实际日志文件名的模式
		rotatelogs.WithLinkName(latestLogFile),    // 生成日志软连接
		rotatelogs.WithMaxAge(30*24*time.Hour),    // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割周期
	)
	if err != nil {
		return nil, err
	}
	logClient.Hooks.Add(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  logWriter,
			logrus.FatalLevel: logWriter,
		},
		&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		},
	))
	return &Logger{logClient}, nil
}
