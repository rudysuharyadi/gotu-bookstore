package auth

import (
	contracts "gotu-bookstore/cmd/gotu-bookstore/contracts/auth"
	"gotu-bookstore/pkg/auth/config"
	"gotu-bookstore/pkg/auth/dto"
	"gotu-bookstore/pkg/utils"

	"gotu-bookstore/cmd/gotu-bookstore/constants"
	"gotu-bookstore/pkg/resfmt/base_error"
)

type LoginService struct {
	context     utils.CommonContext
	userRepo    UserRepoInterface
	authService AuthServiceInterface
	authConfig  config.AuthConfig
}

func NewLoginService(
	context utils.CommonContext,
	userRepo UserRepoInterface,
	authService AuthServiceInterface,
	authConfig config.AuthConfig,
) LoginService {
	return LoginService{
		context:     context,
		userRepo:    userRepo,
		authService: authService,
		authConfig:  authConfig,
	}
}

func (s LoginService) ProcessingLogin(request contracts.LoginRequest) (*contracts.LoginResponse, error) {
	// Get user by email address
	user, err := s.userRepo.GetByEmail(request.Email)
	if err != nil {
		s.context.LogDebug(err)
		return nil, base_error.NewBadRequestError(constants.IC0011)
	}

	// Verify password
	err = s.authService.VerifyPassword(user.Password, request.Password)
	if err != nil {
		s.context.LogDebug(err)
		return nil, base_error.NewBadRequestError(constants.IC0009)
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

	return &contracts.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
