package database

import (
	"context"
	"fmt"
	"github.com/Madfater/AdDeliveryLink/log"

	"gorm.io/gorm/logger"
	"time"
)

type GormLogger struct {
	logLevel      logger.LogLevel
	slowThreshold time.Duration // Threshold for slow queries
}

// NewGormLogger 創建新的 GORM Logger
func NewGormLogger(logLevel logger.LogLevel, slowThreshold time.Duration) *GormLogger {
	return &GormLogger{
		logLevel:      logLevel,
		slowThreshold: slowThreshold,
	}
}

// LogMode 設置日誌模式
func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.logLevel = level
	return &newLogger
}

// Info 記錄 Info 級別日誌
func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.logLevel >= logger.Info {
		log.GetLogger()
	}
}

// Warn 記錄 Warn 級別日誌
func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.logLevel >= logger.Warn {
		log.GetLogger().Warn(fmt.Sprintf(msg, data...), nil)
	}
}

// Error 記錄 Error 級別日誌
func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.logLevel >= logger.Error {
		log.GetLogger().Error(fmt.Sprintf(msg, data...), nil)
	}
}

// Trace 記錄 SQL 語句
func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.logLevel <= 0 {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	fields := map[string]interface{}{
		"sql":      sql,
		"rows":     rows,
		"elapsed":  elapsed.Milliseconds(),
		"duration": elapsed.String(),
	}

	if err != nil {
		fields["error"] = err
		log.GetLogger().Error("SQL execution failed", fields)
		return
	}

	if elapsed > l.slowThreshold {
		log.GetLogger().Warn("Slow SQL detected", fields)
		return
	}

	if l.logLevel >= logger.Info {
		log.GetLogger().Info("SQL executed successfully", fields)
	}
}
