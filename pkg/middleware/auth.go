package middleware

import (
	"gotu-bookstore/pkg/auth/config"
	"gotu-bookstore/pkg/auth/dto"
	"strings"

	"gotu-bookstore/pkg/resfmt/base_error"
	"gotu-bookstore/pkg/resfmt/response_format"

	"gotu-bookstore/cmd/gotu-bookstore/constants"

	"github.com/gin-gonic/gin"
)

type AuthServiceInterface interface {
	VerifyTokenToSessionDTO(config config.BaseConfig, token string, leeway int64) (*dto.SessionDTO, error)
	VerifyToken(config config.BaseConfig, token string, leeway int64) (map[string]interface{}, error)
}

type AuthMiddleware struct {
	authService AuthServiceInterface
	authConfig  config.AuthConfig
}

func NewAuthMiddleware(authService AuthServiceInterface, authConfig config.AuthConfig) AuthMiddleware {
	return AuthMiddleware{
		authService: authService,
		authConfig:  authConfig,
	}
}

func (a AuthMiddleware) UserHandler(leeway int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if len(authHeader) == 0 {
			err := base_error.NewUnauthorizedError(constants.IC0005)
			response := response_format.NewFailure(err)
			c.AbortWithStatusJSON(response.StatusCode, response)
			return
		}

		bearerInfo := strings.Fields(authHeader)
		if len(bearerInfo) < 2 {
			err := base_error.NewUnauthorizedError(constants.IC0004)
			response := response_format.NewFailure(err)
			c.AbortWithStatusJSON(response.StatusCode, response)
			return
		}

		if len(bearerInfo) == 0 {
			err := base_error.NewUnauthorizedError(constants.IC0004)
			response := response_format.NewFailure(err)
			c.AbortWithStatusJSON(response.StatusCode, response)
			return
		}

		accessToken := bearerInfo[1]
		session, err := a.authService.VerifyTokenToSessionDTO(a.authConfig.AccessTokenConfig, accessToken, leeway)
		if err != nil {
			err = base_error.NewUnauthorizedError(constants.IC0004)
			response := response_format.NewFailure(err)
			c.AbortWithStatusJSON(response.StatusCode, response)
			return
		}

		c.Set(constants.AccessTokenContext, accessToken)
		c.Set(constants.SessionDataContext, *session)
		c.Next()
	}
}
