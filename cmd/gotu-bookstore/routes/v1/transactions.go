package v1

import (
	handlers "gotu-bookstore/cmd/gotu-bookstore/handlers/transactions"
	"gotu-bookstore/cmd/gotu-bookstore/routes/dependencies"
	services "gotu-bookstore/cmd/gotu-bookstore/services/transactions"

	"github.com/gin-gonic/gin"
)

func TransactionsRoute(router *gin.RouterGroup, d dependencies.Dependencies) {
	router.GET("transactions/:transaction_id", d.AuthMiddleware.UserHandler(0), func(c *gin.Context) {
		commonContext := d.GetCommonContext(c)
		mainService := services.NewGetTransactionDetailsService(commonContext, d.TransactionsRepository, d.UsersRepository)
		mainHandler := handlers.NewGetTransactionDetailsHandler(commonContext, mainService)
		mainHandler.ProcessingGetTransactionDetails()
	})
	router.GET("transactions", d.AuthMiddleware.UserHandler(0), func(c *gin.Context) {
		commonContext := d.GetCommonContext(c)
		mainService := services.NewGetTransactionsService(commonContext, d.TransactionsRepository, d.UsersRepository)
		mainHandler := handlers.NewGetTransactionsHandler(commonContext, mainService)
		mainHandler.ProcessingGetTransactions()
	})
}
