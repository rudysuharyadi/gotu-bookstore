package middleware

import (
	"net/http"

	"gotu-bookstore/pkg/database"

	"gotu-bookstore/cmd/gotu-bookstore/constants"

	"gotu-bookstore/pkg/logger"

	"github.com/gin-gonic/gin"
)

func statusInList(status int, statusList []int) bool {
	for _, i := range statusList {
		if i == status {
			return true
		}
	}
	return false
}

// DBTransactionMiddleware to wrap the next handlers in a DB transaction
func DBTransactionMiddleware(dbInstance *database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logger.LogInstance
		logContext := logger.NewLogContext().WithField(constants.RequestIDContext, c.Value(constants.RequestIDContext))
		dbTrx := dbInstance.DB.Begin()

		defer func() {
			if r := recover(); r != nil {
				log.ErrorWithContext(logContext, "Rolling back transaction due to panic")
				dbTrx.Rollback()

				// re-throw panic to be handled by gin recovery middleware
				panic(r)
			}
		}()

		c.Set(constants.DBTrxContext, dbTrx)
		c.Next()

		if statusInList(c.Writer.Status(), []int{http.StatusOK, http.StatusCreated}) {
			if err := dbTrx.Commit().Error; err != nil {
				log.ErrorWithContext(logContext, "Unable to commit transaction: ", err)
			}
		} else {
			log.ErrorWithContext(logContext, "Rolling back transaction due to status code: ", c.Writer.Status())
			dbTrx.Rollback()
		}
	}
}
