package database

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func NewMockDB() (*Database, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gormDB, err := gorm.Open(InitPostgreSQLDialectorWithSQLDB(db), &gorm.Config{
		Logger: gormLogger.Default.LogMode(
			gormLogger.LogLevel(1),
		),
	})
	if err != nil {
		return nil, nil, err
	}

	return &Database{DB: *gormDB}, mock, nil
}
