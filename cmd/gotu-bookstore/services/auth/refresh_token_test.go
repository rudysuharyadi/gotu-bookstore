package auth_test

import (
	"gotu-bookstore/cmd/gotu-bookstore/constants"
	contracts "gotu-bookstore/cmd/gotu-bookstore/contracts/auth"
	"gotu-bookstore/cmd/gotu-bookstore/models"
	"gotu-bookstore/cmd/gotu-bookstore/services/auth"
	mockery "gotu-bookstore/mocks/cmd/gotu-bookstore/services/auth"
	"gotu-bookstore/pkg/auth/config"
	"gotu-bookstore/pkg/auth/dto"
	"gotu-bookstore/pkg/utils"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestRefreshTokenServiceSuite(t *testing.T) {
	suite.Run(t, new(RefreshTokenServiceTestSuite))
}

type RefreshTokenServiceTestSuite struct {
	suite.Suite
	fakeUserRepo    *mockery.UserRepoInterface
	fakeAuthService *mockery.AuthServiceInterface
	authConfig      config.AuthConfig
	service         auth.RefreshTokenService
	context         utils.CommonContext
	sessionDTO      dto.SessionDTO
	user            models.Users
}

func (suite *RefreshTokenServiceTestSuite) SetupTest() {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Set(constants.AccessTokenContext, "old-token")
	user := models.Users{
		Id:        uuid.New(),
		Name:      "Rudy Suharyadi",
		Email:     "rudy.suharyadi@gmail.com",
		Password:  "password",
		Status:    string(constants.UserStatusActive),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	sessionDTO := dto.SessionDTO{
		Id:    user.Id.String(),
		Email: user.Email,
	}
	c.Set(constants.SessionDataContext, sessionDTO)
	suite.sessionDTO = sessionDTO
	suite.user = user

	suite.context = utils.NewCommonContext(c, nil, nil)
	suite.fakeUserRepo = mockery.NewUserRepoInterface(suite.T())
	suite.fakeAuthService = mockery.NewAuthServiceInterface(suite.T())
	suite.authConfig = config.AuthConfig{}
	suite.service = auth.NewRefreshTokenService(
		suite.context, suite.fakeUserRepo, suite.fakeAuthService, suite.authConfig)
}

func (suite *RefreshTokenServiceTestSuite) TestRefreshToken() {
	request := contracts.RefreshTokenRequest{
		RefreshToken: "refresh-token",
	}

	refreshSessionDTO := dto.SessionDTO{
		Id:    suite.user.Id.String(),
		Email: suite.user.Email,
	}

	user := models.Users{
		Id:        suite.user.Id,
		Name:      "Rudy Suharyadi",
		Email:     "rudy.suharyadi@gmail.com",
		Password:  "password",
		Status:    string(constants.UserStatusActive),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	suite.fakeAuthService.EXPECT().
		VerifyTokenToSessionDTO(mock.Anything, request.RefreshToken, int64(0)).
		Return(&refreshSessionDTO, nil)
	suite.fakeUserRepo.EXPECT().GetById(refreshSessionDTO.Id).Return(&user, nil)
	suite.fakeAuthService.EXPECT().
		GenerateTokenWithSessionDTO(mock.Anything, suite.sessionDTO).
		Return("token", nil)
	suite.fakeAuthService.EXPECT().
		InvalidateToken(mock.Anything, "old-token").Return(nil)

	response, err := suite.service.ProcessingRefreshToken(request)
	suite.Nil(err)
	suite.NotNil(response)
	suite.NotEmpty(response.AccessToken)
	suite.NotEmpty(response.RefreshToken)
}

func (suite *RefreshTokenServiceTestSuite) TestRefreshToken_ErrorUserIdMismatch() {
	request := contracts.RefreshTokenRequest{
		RefreshToken: "refresh-token",
	}

	refreshSessionDTO := dto.SessionDTO{
		Id:    uuid.New().String(),
		Email: suite.user.Email,
	}

	suite.fakeAuthService.EXPECT().
		VerifyTokenToSessionDTO(mock.Anything, request.RefreshToken, int64(0)).
		Return(&refreshSessionDTO, nil)

	response, err := suite.service.ProcessingRefreshToken(request)
	suite.NotNil(err)
	suite.Nil(response)
}
