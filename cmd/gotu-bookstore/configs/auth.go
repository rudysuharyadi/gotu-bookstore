package configs

import (
	"gotu-bookstore/pkg/auth/config"
	"gotu-bookstore/pkg/viper"
)

var authConfigPrefix = "AUTH_"

func InitAuthConfig() config.AuthConfig {
	return config.AuthConfig{
		AccessTokenConfig: config.BaseConfig{
			SecretKey:  viper.GetStringOrPanic(authConfigPrefix + "ACCESS_TOKEN_SECRET_KEY"),
			Expiration: viper.GetInt64OrPanic(authConfigPrefix + "ACCESS_TOKEN_EXPIRATION"),
		},
		RefreshTokenConfig: config.BaseConfig{
			SecretKey:  viper.GetStringOrPanic(authConfigPrefix + "REFRESH_TOKEN_SECRET_KEY"),
			Expiration: viper.GetInt64OrPanic(authConfigPrefix + "REFRESH_TOKEN_EXPIRATION"),
		},
	}
}

func InitEncryptionConfig() config.EncryptionConfig {
	return config.EncryptionConfig{EncryptionKey: viper.GetStringOrPanic(authConfigPrefix + "ENCRYPTION_KEY")}
}
