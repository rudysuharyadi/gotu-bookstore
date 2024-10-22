package auth

import (
	contracts "gotu-bookstore/cmd/gotu-bookstore/contracts/auth"
	"gotu-bookstore/cmd/gotu-bookstore/models"
	"gotu-bookstore/pkg/auth/config"
	"gotu-bookstore/pkg/auth/dto"
)

type AuthServiceInterface interface {
	VerifyTokenToSessionDTO(config config.BaseConfig, accessToken string, leeway int64) (*dto.SessionDTO, error)
	GenerateTokenWithSessionDTO(config config.BaseConfig, sessionDTO dto.SessionDTO) (string, error)
	InvalidateToken(config config.BaseConfig, accessToken string) error
	HashPassword(password string) (string, error)
	VerifyPassword(hashedPassword string, password string) error
}

type UserRepoInterface interface {
	GetById(id string) (*models.Users, error)
	GetByEmail(email string) (*models.Users, error)
	Create(input models.Users) (*models.Users, error)
}

type UserValidatorInterface interface {
	Validate(request contracts.RegisterRequest) []error
}
