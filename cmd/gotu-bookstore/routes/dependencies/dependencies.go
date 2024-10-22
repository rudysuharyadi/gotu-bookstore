package dependencies

import (
	"gotu-bookstore/cmd/gotu-bookstore/configs"
	"gotu-bookstore/cmd/gotu-bookstore/repositories"
	"gotu-bookstore/pkg/auth/config"
	"gotu-bookstore/pkg/auth/services"
	"gotu-bookstore/pkg/database"
	"gotu-bookstore/pkg/logger"
	"gotu-bookstore/pkg/middleware"
	"gotu-bookstore/pkg/redis"
)

type Dependencies struct {
	RedisInstance    redis.Redis
	AuthConfig       config.AuthConfig
	AuthService      services.AuthService
	AuthMiddleware   middleware.AuthMiddleware
	EncryptionConfig config.EncryptionConfig
	ServerConfig     configs.ServerConfig

	DbInstance             *database.Database
	BooksRepository        repositories.BooksRepository
	ShoppingCartRepository repositories.ShoppingCartsRepository
	TransactionsRepository repositories.TransactionsRepository
	UsersRepository        repositories.UsersRepository

	Log *logger.Log
}

func NewDependencies() Dependencies {
	// _ := logger.LogInstance
	redisInstance := redis.Instance
	authConfig := configs.ConfigInstance.Auth
	authService := services.NewAuthService(redisInstance)
	authMiddleware := middleware.NewAuthMiddleware(authService, authConfig)
	encryptionConfig := configs.ConfigInstance.Encryption
	serverConfig := configs.ConfigInstance.Server

	dbInstance := database.DBInstance
	booksRepository := repositories.NewBooksRepository(dbInstance)
	shoppingCartRepository := repositories.NewShoppingCartsRepository(dbInstance)
	transactionsRepository := repositories.NewTransactionsRepository(dbInstance)
	usersRepository := repositories.NewUsersRepository(dbInstance)

	return Dependencies{
		RedisInstance:          redisInstance,
		AuthConfig:             authConfig,
		AuthService:            authService,
		AuthMiddleware:         authMiddleware,
		EncryptionConfig:       encryptionConfig,
		ServerConfig:           serverConfig,
		DbInstance:             dbInstance,
		BooksRepository:        booksRepository,
		ShoppingCartRepository: shoppingCartRepository,
		TransactionsRepository: transactionsRepository,
		UsersRepository:        usersRepository,
	}
}
