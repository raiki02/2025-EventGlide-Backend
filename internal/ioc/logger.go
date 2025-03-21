package ioc

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func Newlogger() *zap.Logger {
	eConf := zap.NewProductionEncoderConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(eConf),
		zapcore.AddSync(zapcore.AddSync(os.Stdout)),
		zapcore.InfoLevel,
	)
	return zap.New(core)
}
