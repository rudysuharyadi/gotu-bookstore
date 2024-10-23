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

func TestRegisterServiceSuite(t *testing.T) {
	suite.Run(t, new(RegisterServiceTestSuite))
}

type RegisterServiceTestSuite struct {
	suite.Suite
	fakeUserRepo      *mockery.UserRepoInterface
	fakeUserValidator *mockery.UserValidatorInterface
	fakeAuthService   *mockery.AuthServiceInterface
	authConfig        config.AuthConfig
	service           auth.RegisterService
	context           utils.CommonContext
}

func (suite *RegisterServiceTestSuite) SetupTest() {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Set(constants.AccessTokenContext, "token")
	suite.context = utils.NewCommonContext(c, nil, nil)
	suite.fakeUserRepo = mockery.NewUserRepoInterface(suite.T())
	suite.fakeUserValidator = mockery.NewUserValidatorInterface(suite.T())
	suite.fakeAuthService = mockery.NewAuthServiceInterface(suite.T())
	suite.authConfig = config.AuthConfig{}
	suite.service = auth.NewRegisterService(
		suite.context,
		suite.fakeUserValidator,
		suite.fakeUserRepo,
		suite.fakeAuthService,
		suite.authConfig,
	)
}

func (suite *RegisterServiceTestSuite) TestRegister() {
	request := contracts.RegisterRequest{
		Password:        "password",
		ConfirmPassword: "password",
		Name:            "name",
		Email:           "email",
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

	suite.fakeUserValidator.EXPECT().Validate(request).Return(nil)
	suite.fakeAuthService.EXPECT().HashPassword(request.Password).Return("hashed", nil)
	suite.fakeUserRepo.EXPECT().Create(mock.Anything).Return(&user, nil)

	response, err := suite.service.ProcessingRegister(request)

	suite.Nil(err)
	suite.Empty(response)
}
