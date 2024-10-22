package configs

import (
	"gotu-bookstore/pkg/viper"
)

type ServerConfig struct {
	AppMode          string
	AppBaseUrl       string
	AppHost          string
	AppPort          int
	GlobalTimeout    int
	GlobalRetryCount int
}

var serverConfigPrefix = "SERVER_"

func InitServerConfig() ServerConfig {
	return ServerConfig{
		AppMode:          viper.GetStringOrPanic(serverConfigPrefix + "APP_MODE"),
		AppHost:          viper.GetStringOrPanic(serverConfigPrefix + "APP_HOST"),
		AppPort:          viper.GetIntOrPanic(serverConfigPrefix + "APP_PORT"),
		GlobalTimeout:    viper.GetIntOrPanic(serverConfigPrefix + "GLOBAL_TIMEOUT"),
		GlobalRetryCount: viper.GetIntOrPanic(serverConfigPrefix + "GLOBAL_RETRY_COUNT"),
	}
}
