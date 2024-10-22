package redis

type RedisConfig struct {
	Host                   string
	Port                   int
	DefaultCacheExpiration int
	Username               string
	Password               string
}
