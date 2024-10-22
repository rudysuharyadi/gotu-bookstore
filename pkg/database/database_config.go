package database

// DatabaseConfig contains database configurations
type DatabaseConfig struct {
	DbName          string
	Host            string
	Port            int
	Username        string
	Password        string
	MaxIddleConn    int
	MaxOpenConn     int
	ConnMaxLifetime int
	LogMode         int
}
