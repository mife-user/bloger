package exc

import (
	"mifer/pkg/logger"
	"os"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// TestMain 测试主函数
func TestMain(m *testing.M) {
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:    "time",
		LevelKey:   "level",
		MessageKey: "msg",
		LineEnding: zapcore.DefaultLineEnding,
	}

	encoder := zapcore.NewConsoleEncoder(encoderCfg)
	core := zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel)
	testLogger := zap.New(core)

	logger.SetLogger(testLogger)
	zap.ReplaceGlobals(testLogger)

	os.Exit(m.Run())
}
