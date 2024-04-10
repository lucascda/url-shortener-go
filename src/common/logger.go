package common

import "go.uber.org/zap"

var Logger *zap.SugaredLogger

func InitLogger() *zap.SugaredLogger {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	return logger.Sugar()
}
