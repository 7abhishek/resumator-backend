package logger

import "go.uber.org/zap"

var loggerSugar *zap.SugaredLogger

func GetLogger() *zap.SugaredLogger {
	if loggerSugar == nil {
		logger, _ := zap.NewProduction()
		loggerSugar = logger.Sugar()
	}
	return loggerSugar
}
