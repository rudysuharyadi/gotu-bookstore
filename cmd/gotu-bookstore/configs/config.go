package configs

import (
	"gotu-bookstore/pkg/adapter/api"
	"gotu-bookstore/pkg/auth/config"
	"gotu-bookstore/pkg/database"
	"gotu-bookstore/pkg/logger"
	"gotu-bookstore/pkg/redis"

	"github.com/spf13/viper"
)

var ConfigInstance AppConfig

type AppConfig struct {
	Server         ServerConfig
	Log            logger.LoggerConfig
	DatabaseConfig database.DatabaseConfig
	RedisConfig    redis.RedisConfig
	Auth           config.AuthConfig
	Encryption     config.EncryptionConfig
	MailgunConfig  api.MailgunConfig
}

// InitConfig loads configs from files
func InitConfig(fileName string, log logger.Log) AppConfig {
	viper.AutomaticEnv()
	viper.SetConfigType("json")
	viper.SetConfigName(fileName)
	viper.AddConfigPath("./resources")

	if err := viper.ReadInConfig(); err != nil {
		log.Errorf("failed to read the configuration file: %s", err.Error())
	}

	ConfigInstance = AppConfig{
		Server:         InitServerConfig(),
		Log:            InitLogConfig(),
		DatabaseConfig: InitDatabaseConfig(),
		RedisConfig:    InitRedisConfig(),
		Auth:           InitAuthConfig(),
		Encryption:     InitEncryptionConfig(),
	}

	log.Infof("configuration successfully loaded")
	return ConfigInstance
}
