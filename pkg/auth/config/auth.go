package config

type AuthConfig struct {
	AccessTokenConfig  BaseConfig
	RefreshTokenConfig BaseConfig
}

type BaseConfig struct {
	SecretKey  string
	Expiration int64
}

type EncryptionConfig struct {
	EncryptionKey string
}
