package log

import (
	"sync"
)

var (
	once     sync.Once
	instance ILogger
)

func GetLogger() ILogger {
	once.Do(func() {
		fileLogger := NewLogrusLogger("./logs/app.log")
		instance = fileLogger
	})
	return instance
}

func SetLogger(customLogger ILogger) {
	once.Do(func() {
		instance = customLogger
	})
}
