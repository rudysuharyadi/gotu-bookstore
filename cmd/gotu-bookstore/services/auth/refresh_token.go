package auth

import (
	"gotu-bookstore/cmd/gotu-bookstore/constants"
	contracts "gotu-bookstore/cmd/gotu-bookstore/contracts/auth"
	"gotu-bookstore/pkg/auth/config"
	"gotu-bookstore/pkg/auth/dto"
	"gotu-bookstore/pkg/resfmt/base_error"
	"gotu-bookstore/pkg/utils"
)

type RefreshTokenService struct {
	context     utils.CommonContext
	userRepo    UserRepoInterface
	authService AuthServiceInterface
	authConfig  config.AuthConfig
}

func NewRefreshTokenService(
	context utils.CommonContext,
	userRepo UserRepoInterface,
	authService AuthServiceInterface,
	authConfig config.AuthConfig,
) RefreshTokenService {
	return RefreshTokenService{
		context:     context,
		userRepo:    userRepo,
		authService: authService,
		authConfig:  authConfig,
	}
}

func (s RefreshTokenService) ProcessingRefreshToken(request contracts.RefreshTokenRequest) (*contracts.RefreshTokenResponse, error) {
	// Get old access token
	oldAccessToken, err := s.context.GetAccessToken()
	if err != nil {
		s.context.LogError(err)
		return nil, base_error.NewUnauthorizedError(constants.IC0006)
	}

	oldSession, err := s.context.GetSession()
	if err != nil {
		s.context.LogError(err)
		return nil, base_error.NewUnauthorizedError(constants.IC0006)
	}

	refreshSession, err := s.authService.VerifyTokenToSessionDTO(s.authConfig.RefreshTokenConfig, request.RefreshToken, 0)
	if err != nil {
		s.context.LogDebug(err)
		return nil, base_error.NewUnauthorizedError(constants.IC0006)
	}

	if refreshSession.Id != oldSession.Id {
		err = base_error.NewUnauthorizedError(constants.IC0006)
		s.context.LogError(err)
		return nil, err
	}

	// Get user by Id
	user, err := s.userRepo.GetById(refreshSession.Id)
	if err != nil {
		s.context.LogDebug(err)
		return nil, base_error.NewInternalError(constants.IC0011)
	}

	// Generate access token and refresh token
	sessionDTO := dto.SessionDTO{
		Email: user.Email,
		Id:    user.Id.String(),
	}

	accessToken, err := s.authService.GenerateTokenWithSessionDTO(s.authConfig.AccessTokenConfig, sessionDTO)
	if err != nil {
		s.context.LogDebug(err)
		return nil, base_error.NewInternalError(constants.IC0012)
	}

	refreshToken, err := s.authService.GenerateTokenWithSessionDTO(s.authConfig.RefreshTokenConfig, sessionDTO)
	if err != nil {
		s.context.LogDebug(err)
		return nil, base_error.NewInternalError(constants.IC0012)
	}

	// Invalidate access token by adding to banned token list.
	err = s.authService.InvalidateToken(s.authConfig.AccessTokenConfig, oldAccessToken)
	if err != nil {
		s.context.LogError(err)
		return nil, base_error.NewInternalError(constants.IC0013)
	}

	return &contracts.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
