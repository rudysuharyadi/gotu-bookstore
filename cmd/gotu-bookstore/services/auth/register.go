package auth

import (
	"gotu-bookstore/cmd/gotu-bookstore/constants"
	contracts "gotu-bookstore/cmd/gotu-bookstore/contracts/auth"
	"gotu-bookstore/cmd/gotu-bookstore/models"
	"gotu-bookstore/pkg/auth/config"
	"gotu-bookstore/pkg/resfmt/base_error"
	"gotu-bookstore/pkg/utils"
	"strings"
	"time"

	"github.com/google/uuid"
)

type RegisterService struct {
	context       utils.CommonContext
	userValidator UserValidatorInterface
	userRepo      UserRepoInterface
	authService   AuthServiceInterface
	authConfig    config.AuthConfig
}

func NewRegisterService(
	context utils.CommonContext,
	userValidator UserValidatorInterface,
	userRepo UserRepoInterface,
	authService AuthServiceInterface,
	authConfig config.AuthConfig,
) RegisterService {
	return RegisterService{
		context:       context,
		userValidator: userValidator,
		userRepo:      userRepo,
		authService:   authService,
		authConfig:    authConfig,
	}
}

func (s RegisterService) ProcessingRegister(request contracts.RegisterRequest) (*contracts.RegisterResponse, error) {
	errs := s.userValidator.Validate(request)
	if len(errs) > 0 {
		s.context.LogError(errs[0])
		return nil, base_error.NewBadRequestErrorWithCause(errs[0], constants.IC0009)
	}

	_, err := s.CreateUser(request.Name, request.Email, request.Password)
	if err != nil {
		s.context.LogError(err)
		return nil, base_error.NewInternalError(constants.IC0010)
	}

	return &contracts.RegisterResponse{}, nil
}

func (s RegisterService) CreateUser(name, email, password string) (*models.Users, error) {
	// Hash the password
	hashedPassword, err := s.authService.HashPassword(password)
	if err != nil {
		return nil, err
	}

	userId, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.Create(models.Users{
		Id:        userId,
		Name:      name,
		Email:     strings.ToLower(email),
		Password:  hashedPassword,
		Status:    string(constants.UserStatusActive),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}
