package v1

import (
	handlers "gotu-bookstore/cmd/gotu-bookstore/handlers/auth"
	"gotu-bookstore/cmd/gotu-bookstore/routes/dependencies"
	services "gotu-bookstore/cmd/gotu-bookstore/services/auth"

	"github.com/gin-gonic/gin"
)

func AuthRoute(router *gin.RouterGroup, d dependencies.Dependencies) {
	router.POST("login", func(c *gin.Context) {
		commonContext := d.GetCommonContext(c)
		mainService := services.NewLoginService(commonContext, d.UsersRepository, d.AuthService, d.AuthConfig)
		mainHandler := handlers.NewLoginHandler(commonContext, mainService)
		mainHandler.ProcessingLogin()
	})
	router.POST("logout", d.AuthMiddleware.UserHandler(0), func(c *gin.Context) {
		commonContext := d.GetCommonContext(c)
		mainService := services.NewLogoutService(commonContext, d.AuthService, d.AuthConfig)
		mainHandler := handlers.NewLogoutHandler(commonContext, mainService)
		mainHandler.ProcessingLogout()
	})
	refreshTokenLeeway := d.AuthConfig.RefreshTokenConfig.Expiration - d.AuthConfig.AccessTokenConfig.Expiration
	router.POST("refresh-token",
		d.AuthMiddleware.UserHandler(refreshTokenLeeway),
		func(c *gin.Context) {
			commonContext := d.GetCommonContext(c)
			mainService := services.NewRefreshTokenService(commonContext, d.UsersRepository, d.AuthService, d.AuthConfig)
			mainHandler := handlers.NewRefreshTokenHandler(commonContext, mainService)
			mainHandler.ProcessingRefreshToken()
		},
	)
	router.POST("register", func(c *gin.Context) {
		commonContext := d.GetCommonContext(c)
		userValidator := services.NewUsersValidator()
		mainService := services.NewRegisterService(commonContext, userValidator, d.UsersRepository, d.AuthService, d.AuthConfig)
		mainHandler := handlers.NewRegisterHandler(commonContext, mainService)
		mainHandler.ProcessingRegister()
	})
}
