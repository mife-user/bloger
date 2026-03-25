package logger

import "go.uber.org/zap"

// SetLogger 设置logger实例（用于测试）
func SetLogger(logger *zap.Logger) {
	loggerInstance = logger
}
