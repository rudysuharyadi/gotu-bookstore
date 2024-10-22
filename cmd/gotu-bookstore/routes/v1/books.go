package v1

import (
	handlers "gotu-bookstore/cmd/gotu-bookstore/handlers/books"
	"gotu-bookstore/cmd/gotu-bookstore/routes/dependencies"
	services "gotu-bookstore/cmd/gotu-bookstore/services/books"

	"github.com/gin-gonic/gin"
)

func BooksRoute(router *gin.RouterGroup, d dependencies.Dependencies) {
	router.GET("books", func(c *gin.Context) {
		commonContext := d.GetCommonContext(c)
		mainService := services.NewGetBooksService(commonContext, d.BooksRepository)
		mainHandler := handlers.NewGetBooksHandler(commonContext, mainService)
		mainHandler.ProcessingGetBooks()
	})
	router.GET("books/:book_id", func(c *gin.Context) {
		commonContext := d.GetCommonContext(c)
		mainService := services.NewGetBookDetailsService(commonContext, d.BooksRepository)
		mainHandler := handlers.NewGetBookDetailsHandler(commonContext, mainService)
		mainHandler.ProcessingGetBookDetails()
	})
}
