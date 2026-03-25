package gittool

import (
	"bloger/pkg/logger"
	"os"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// TestMain 测试主函数，初始化logger
func TestMain(m *testing.M) {
	// 创建测试用的zap logger
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:    "time",
		LevelKey:   "level",
		MessageKey: "msg",
		LineEnding: zapcore.DefaultLineEnding,
	}

	encoder := zapcore.NewConsoleEncoder(encoderCfg)
	core := zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel)
	testLogger := zap.New(core)

	// 设置logger包的实例
	logger.SetLogger(testLogger)
	zap.ReplaceGlobals(testLogger)

	os.Exit(m.Run())
}
