package routes

import (
	"gotu-bookstore/cmd/gotu-bookstore/routes/dependencies"
	v1 "gotu-bookstore/cmd/gotu-bookstore/routes/v1"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) *gin.Engine {
	d := dependencies.NewDependencies()

	accountV1Path := r.Group("/account-service/v1")
	{
		v1.AuthRoute(accountV1Path, d)
	}

	productV1Path := r.Group("/product-service/v1")
	{
		v1.BooksRoute(productV1Path, d)
	}

	orderV1Path := r.Group("order-service/v1")
	{
		v1.ShoppingCartRoute(orderV1Path, d)
		v1.TransactionsRoute(orderV1Path, d)
	}

	return r
}
