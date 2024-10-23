package auth_test

import (
	"gotu-bookstore/cmd/gotu-bookstore/constants"
	contracts "gotu-bookstore/cmd/gotu-bookstore/contracts/auth"
	"gotu-bookstore/cmd/gotu-bookstore/models"
	"gotu-bookstore/cmd/gotu-bookstore/services/auth"
	mockery "gotu-bookstore/mocks/cmd/gotu-bookstore/services/auth"
	"gotu-bookstore/pkg/auth/config"
	"gotu-bookstore/pkg/utils"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestLoginServiceSuite(t *testing.T) {
	suite.Run(t, new(LoginServiceTestSuite))
}

type LoginServiceTestSuite struct {
	suite.Suite
	fakeUserRepo    *mockery.UserRepoInterface
	fakeAuthService *mockery.AuthServiceInterface
	authConfig      config.AuthConfig
	service         auth.LoginService
	context         utils.CommonContext
}

func (suite *LoginServiceTestSuite) SetupTest() {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	suite.context = utils.NewCommonContext(c, nil, nil)
	suite.fakeUserRepo = mockery.NewUserRepoInterface(suite.T())
	suite.fakeAuthService = mockery.NewAuthServiceInterface(suite.T())
	suite.authConfig = config.AuthConfig{}
	suite.service = auth.NewLoginService(suite.context, suite.fakeUserRepo, suite.fakeAuthService, suite.authConfig)
}

func (suite *LoginServiceTestSuite) TestLogin() {
	request := contracts.LoginRequest{
		Email:    "rudy.suharyadi@gmail.com",
		Password: "password",
	}

	user := models.Users{
		Id:        uuid.New(),
		Name:      "Rudy Suharyadi",
		Email:     "rudy.suharyadi@gmail.com",
		Password:  "password",
		Status:    string(constants.UserStatusActive),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	suite.fakeUserRepo.EXPECT().GetByEmail(request.Email).Return(&user, nil)
	suite.fakeAuthService.EXPECT().VerifyPassword(user.Password, request.Password).Return(nil)
	suite.fakeAuthService.EXPECT().GenerateTokenWithSessionDTO(mock.Anything, mock.Anything).Return("token", nil)

	response, err := suite.service.ProcessingLogin(request)
	suite.Nil(err)
	suite.NotNil(response)
	suite.NotEmpty(response.AccessToken)
	suite.NotEmpty(response.RefreshToken)
}
