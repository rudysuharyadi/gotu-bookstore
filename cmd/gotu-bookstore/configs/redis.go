package configs

import (
	"gotu-bookstore/pkg/redis"
	"gotu-bookstore/pkg/viper"
)

var redisConfigPrefix = "REDIS_"

func InitRedisConfig() redis.RedisConfig {
	return redis.RedisConfig{
		Host:                   viper.GetStringOrPanic(redisConfigPrefix + "HOST"),
		Port:                   viper.GetIntOrPanic(redisConfigPrefix + "PORT"),
		DefaultCacheExpiration: viper.GetIntOrPanic(redisConfigPrefix + "DEFAULT_CACHE_EXPIRATION"),
		Username:               viper.GetStringOrPanic(redisConfigPrefix + "USERNAME"),
		Password:               viper.GetStringOrPanic(redisConfigPrefix + "PASSWORD"),
	}
}
