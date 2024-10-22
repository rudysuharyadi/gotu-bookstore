package main

import (
	"embed"
	"os"

	"gotu-bookstore/cmd/gotu-bookstore/configs"
	"gotu-bookstore/cmd/gotu-bookstore/constants"

	"gotu-bookstore/pkg/database"
	"gotu-bookstore/pkg/logger"
	"gotu-bookstore/pkg/middleware"
	"gotu-bookstore/pkg/redis"
	"gotu-bookstore/pkg/resfmt/base_error"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	// initialize logger instance with default configuration
	log := logger.NewLog()

	// initialize configuration data
	appConfig := configs.InitConfig("application", log)

	// update logger instance based on configuration
	err := log.UpdateLevel(appConfig.Log.Level)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	// init http client
	initHTTPClient()

	// initialize database connection
	dialector := database.InitPostgreSQLDialector(appConfig.DatabaseConfig)
	dbInstance, err := database.NewDatabase(dialector, appConfig.DatabaseConfig)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	log.Infof("Database Initialized")

	// initialize redis connection
	redis.NewRedis(appConfig.RedisConfig)
	log.Infof("Redis Initialized")

	err = dbInstance.EmbedMigrations(dialector, embedMigrations, "migrations")
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	// Set base_error ErrorMessage
	base_error.ErrorMessages = constants.ErrorMessages

	// initialize gin framework & configuration
	NewEngineBuilder(appConfig).
		SetProxy().
		RegisterMiddleware(middleware.RequestID()).
		RegisterMiddleware(middleware.BeforeAfterRequest()).
		RegisterMiddleware(middleware.CustomRecoveryMiddleware()).
		RegisterMiddleware(middleware.DBTransactionMiddleware(dbInstance)).
		RegisterNoRoute(middleware.NoRouteMiddleware()).
		InitRoutes().
		Run()
}

func initHTTPClient() {

}
