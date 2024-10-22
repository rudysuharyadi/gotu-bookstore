package auth

import (
	"gotu-bookstore/cmd/gotu-bookstore/constants"
	"gotu-bookstore/pkg/auth/config"
	"gotu-bookstore/pkg/resfmt/base_error"
	"gotu-bookstore/pkg/utils"
)

type LogoutService struct {
	context     utils.CommonContext
	authService AuthServiceInterface
	authConfig  config.AuthConfig
}

func NewLogoutService(
	context utils.CommonContext,
	authService AuthServiceInterface,
	authConfig config.AuthConfig,
) LogoutService {
	return LogoutService{
		context:     context,
		authService: authService,
		authConfig:  authConfig,
	}
}

func (s LogoutService) ProcessingLogout() error {
	// Get Access token
	accessToken, err := s.context.GetAccessToken()
	if err != nil {
		s.context.LogError(err)
		return base_error.NewUnauthorizedError(constants.IC0006)
	}

	// Invalidate access token by adding to banned token list.
	err = s.authService.InvalidateToken(s.authConfig.AccessTokenConfig, accessToken)
	if err != nil {
		s.context.LogError(err)
		return base_error.NewInternalError(constants.IC0013)
	}

	return nil
}
