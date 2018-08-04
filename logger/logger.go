package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/kensay98/vindecoder/config"
)

var (
	logger *logrus.Logger
	logLevels = map[string]logrus.Level {
		"INFO": logrus.InfoLevel,
		"DEBUG": logrus.DebugLevel,
		"ERROR": logrus.ErrorLevel,
	}
)

func GetLogger() *logrus.Logger{
	if logger != nil {
		return logger
	}

	logger = logrus.New()
	logLevel, ok := logLevels[config.GetConfig().LogLevel]
	if ok != true {
		panic("Incorrect log level is set")
	}

	logger.SetLevel(logLevel)
	return logger
}
