// internal/logger/logger.go
package logger

import (
	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

func InitLogger() {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	Log = zapLogger.Sugar()
}

