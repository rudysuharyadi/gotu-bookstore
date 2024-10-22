package dependencies

import (
	"gotu-bookstore/cmd/gotu-bookstore/constants"
	"gotu-bookstore/pkg/logger"
	"gotu-bookstore/pkg/utils"

	"github.com/gin-gonic/gin"
)

func (d Dependencies) GetCommonContext(c *gin.Context) utils.CommonContext {
	logContext := logger.NewLogContext().WithField(constants.RequestIDContext, c.Value(constants.RequestIDContext))
	return utils.CommonContext{
		GinContext:  c,
		LogInstance: d.Log,
		LogContext:  &logContext,
	}
}
