package logger

import (
	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var L *zap.SugaredLogger

func init() {
	zapConfig := zap.NewDevelopmentEncoderConfig()
	zapConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zapConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	core := zapcore.NewCore(zapcore.NewConsoleEncoder(zapConfig),
		zapcore.AddSync(colorable.NewColorableStdout()),
		zapcore.DebugLevel,
	)

	L = zap.New(core).Sugar()
}
