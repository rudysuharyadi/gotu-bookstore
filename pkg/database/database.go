package database

import (
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

var DBInstance *Database

type Database struct {
	DB gorm.DB
}

func NewDatabase(databaseDialector gorm.Dialector, config DatabaseConfig) (*Database, error) {
	logMode := gormLogger.Silent
	if config.LogMode > 0 {
		logMode = gormLogger.LogLevel(config.LogMode)
	}

	db, err := gorm.Open(databaseDialector, &gorm.Config{
		Logger: gormLogger.Default.LogMode(logMode),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(config.MaxIddleConn)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(config.MaxOpenConn)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime) * time.Hour)

	DBInstance = &Database{DB: *db}
	return DBInstance, nil
}
