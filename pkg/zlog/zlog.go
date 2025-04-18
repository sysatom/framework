package zlog

import (
	"go.uber.org/zap"
)

func NewZlog() *zap.Logger {
	logger, _ := zap.NewProduction()
	// defer logger.Sync() // flushes buffer, if any
	return logger
}
