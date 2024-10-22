package middleware

import (
	"bytes"
	"fmt"
	"gotu-bookstore/cmd/gotu-bookstore/constants"
	"io"
	"time"

	"gotu-bookstore/pkg/logger"

	"github.com/gin-gonic/gin"
)

func BeforeAfterRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log := logger.LogInstance
		logContext := logger.NewLogContext().WithField(constants.RequestIDContext, ctx.Value(constants.RequestIDContext))

		// fullPath should be not empty. However, there might be some cases where fullPath is empty, like no route.
		fullPath := ctx.FullPath()
		if len(fullPath) == 0 {
			fullPath = ctx.Request.RequestURI
		}

		// get request body
		var bodyBytes []byte
		if ctx.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(ctx.Request.Body)
		}
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		if len(bodyBytes) > 0 {
			log.DebugfWithContext(logContext, "Request body:\n%s", string(bodyBytes))
		}

		start := time.Now()
		ctx.Next()
		elapsed := time.Since(start)

		for idx, err := range ctx.Errors {
			log.ErrorWithContext(logContext, fmt.Sprintf("#%d %s", idx, err.Error()))
		}
		log.InfofWithContext(logContext, "%s %s | %d | %s | %.6fms", ctx.Request.Method, fullPath, ctx.Writer.Status(), ctx.ClientIP(), float64(elapsed.Nanoseconds())/1e6)
	}
}
