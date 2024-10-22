package middleware

import (
	"gotu-bookstore/cmd/gotu-bookstore/constants"

	"gotu-bookstore/pkg/resfmt/base_error"
	"gotu-bookstore/pkg/resfmt/response_format"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestID, err := uuid.NewV7()
		if err != nil {
			brerr := base_error.NewInternalError(constants.IC0002)
			response := response_format.NewFailure(brerr)
			ctx.Error(brerr)
			ctx.AbortWithStatusJSON(response.StatusCode, response)
		}

		ctx.Set(constants.RequestIDContext, requestID.String())
		ctx.Writer.Header().Set("x-request-id", requestID.String())

		ctx.Next()
	}
}
