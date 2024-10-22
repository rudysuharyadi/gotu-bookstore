package middleware

import (
	"fmt"
	"gotu-bookstore/cmd/gotu-bookstore/constants"

	"gotu-bookstore/pkg/logger"
	"gotu-bookstore/pkg/resfmt/base_error"
	"gotu-bookstore/pkg/resfmt/response_format"

	"github.com/gin-gonic/gin"
)

func CustomRecoveryMiddleware() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		log := logger.LogInstance
		logContext := logger.NewLogContext().WithField(constants.RequestIDContext, c.Value(constants.RequestIDContext))
		log.DebugWithContext(logContext, fmt.Sprintf("panic : %v", recovered))

		err := base_error.NewInternalError(constants.IC0007)
		response := response_format.NewFailure(err)
		c.Error(err)
		c.AbortWithStatusJSON(response.StatusCode, response)
	})
}
