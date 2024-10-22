package database

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GetDatabaseConnectionString takes DB configuration and returns a DB connection string
func getPostgreSQLDatabaseConnectionString(config DatabaseConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC", config.Host, config.Username, config.Password, config.DbName, config.Port)
}

func InitPostgreSQLDialector(config DatabaseConfig) gorm.Dialector {
	return postgres.Open(getPostgreSQLDatabaseConnectionString(config))
}

func InitPostgreSQLDialectorWithSQLDB(sqlDB *sql.DB) gorm.Dialector {
	return postgres.New(postgres.Config{
		Conn:       sqlDB,
		DriverName: "postgres",
	})
}
