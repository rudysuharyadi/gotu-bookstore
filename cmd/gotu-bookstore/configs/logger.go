package configs

import (
	"gotu-bookstore/pkg/logger"
	"gotu-bookstore/pkg/viper"
)

var logConfigPrefix = "LOG_"

func InitLogConfig() logger.LoggerConfig {
	return logger.LoggerConfig{
		Level: viper.GetStringOrPanic(logConfigPrefix + "LEVEL"),
	}
}
