package initialize

import (
	"os"

	"github.com/leedev/go-rest-ddd/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() *zap.Logger {
	cfg := zap.NewProductionEncoderConfig()

	cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	cfg.EncodeTime = zapcore.ISO8601TimeEncoder

	cfg.TimeKey = "time"

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg),
		zapcore.AddSync(os.Stdout),
		zap.DebugLevel,
	)

	global.Logger = zap.New(core)
	return global.Logger
}
