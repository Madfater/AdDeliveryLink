package log

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LogrusLogger struct {
	logger *logrus.Logger
}

func NewLogrusLogger(logFile string) *LogrusLogger {
	l := logrus.New()

	l.Out = &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    10,   // 文件最大尺寸 (MB)
		MaxBackups: 5,    // 保留的舊文件數量
		MaxAge:     30,   // 最長保留天數
		Compress:   true, // 是否壓縮舊文件
	}

	l.SetFormatter(&logrus.JSONFormatter{})

	return &LogrusLogger{logger: l}
}

func (l *LogrusLogger) Info(message string, fields map[string]interface{}) {
	l.logger.WithFields(logrus.Fields(fields)).Info(message)
}

func (l *LogrusLogger) Error(message string, fields map[string]interface{}) {
	l.logger.WithFields(logrus.Fields(fields)).Error(message)
}

func (l *LogrusLogger) Debug(message string, fields map[string]interface{}) {
	l.logger.WithFields(logrus.Fields(fields)).Debug(message)
}

func (l *LogrusLogger) Warn(message string, fields map[string]interface{}) {
	l.logger.WithFields(logrus.Fields(fields)).Warn(message)
}
