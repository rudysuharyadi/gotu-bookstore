package configs

import (
	"gotu-bookstore/pkg/database"
	"gotu-bookstore/pkg/viper"
)

var databaseConfigPrefix = "DATABASE_"

func InitDatabaseConfig() database.DatabaseConfig {
	return database.DatabaseConfig{
		DbName:          viper.GetStringOrPanic(databaseConfigPrefix + "DB_NAME"),
		Host:            viper.GetStringOrPanic(databaseConfigPrefix + "HOST"),
		Port:            viper.GetIntOrPanic(databaseConfigPrefix + "PORT"),
		Username:        viper.GetStringOrPanic(databaseConfigPrefix + "USERNAME"),
		Password:        viper.GetStringOrPanic(databaseConfigPrefix + "PASSWORD"),
		MaxIddleConn:    viper.GetIntOrPanic(databaseConfigPrefix + "MAX_IDDLE_CONN"),
		MaxOpenConn:     viper.GetIntOrPanic(databaseConfigPrefix + "MAX_OPEN_CONN"),
		ConnMaxLifetime: viper.GetIntOrPanic(databaseConfigPrefix + "CONN_MAX_LIFETIME"),
		LogMode:         viper.GetIntOrPanic(databaseConfigPrefix + "LOG_MODE"),
	}
}
