package v1

import (
	handlers "gotu-bookstore/cmd/gotu-bookstore/handlers/shopping_cart"
	"gotu-bookstore/cmd/gotu-bookstore/routes/dependencies"
	services "gotu-bookstore/cmd/gotu-bookstore/services/shopping_cart"

	"github.com/gin-gonic/gin"
)

func ShoppingCartRoute(router *gin.RouterGroup, d dependencies.Dependencies) {
	router.POST("shopping-cart", d.AuthMiddleware.UserHandler(0), func(c *gin.Context) {
		commonContext := d.GetCommonContext(c)
		mainService := services.NewPostShoppingCartService(commonContext, d.ShoppingCartRepository, d.UsersRepository, d.BooksRepository)
		mainHandler := handlers.NewPostShoppingCartHandler(commonContext, mainService)
		mainHandler.ProcessingPostShoppingCart()
	})
	router.GET("shopping-cart", d.AuthMiddleware.UserHandler(0), func(c *gin.Context) {
		commonContext := d.GetCommonContext(c)
		mainService := services.NewGetShoppingCartService(commonContext, d.ShoppingCartRepository, d.UsersRepository)
		mainHandler := handlers.NewGetShoppingCartHandler(commonContext, mainService)
		mainHandler.ProcessingGetShoppingCart()
	})
	router.POST("shopping-cart/checkout", d.AuthMiddleware.UserHandler(0), func(c *gin.Context) {
		commonContext := d.GetCommonContext(c)
		mainService := services.NewPostShoppingCartCheckoutService(commonContext, d.ShoppingCartRepository, d.UsersRepository, d.TransactionsRepository)
		mainHandler := handlers.NewPostShoppingCartCheckoutHandler(commonContext, mainService)
		mainHandler.ProcessingPostShoppingCartCheckout()
	})
	router.POST("shopping-cart/clear", d.AuthMiddleware.UserHandler(0), func(c *gin.Context) {
		commonContext := d.GetCommonContext(c)
		mainService := services.NewPostShoppingCartClearService(commonContext, d.ShoppingCartRepository, d.UsersRepository)
		mainHandler := handlers.NewPostShoppingCartClearHandler(commonContext, mainService)
		mainHandler.ProcessingPostShoppingCartClear()
	})
}
