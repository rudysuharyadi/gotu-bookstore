package middleware

import (
	"gotu-bookstore/cmd/gotu-bookstore/constants"

	"gotu-bookstore/pkg/resfmt/base_error"
	"gotu-bookstore/pkg/resfmt/response_format"

	"github.com/gin-gonic/gin"
)

func NoRouteMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		brerr := base_error.NewNotFoundError(constants.IC0008)
		response := response_format.NewFailure(brerr)
		c.Error(brerr)
		c.AbortWithStatusJSON(response.StatusCode, response)
	}
}
