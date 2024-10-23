package auth_test

import (
	"gotu-bookstore/cmd/gotu-bookstore/constants"
	"gotu-bookstore/cmd/gotu-bookstore/services/auth"
	mockery "gotu-bookstore/mocks/cmd/gotu-bookstore/services/auth"
	"gotu-bookstore/pkg/auth/config"
	"gotu-bookstore/pkg/utils"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestLogoutServiceSuite(t *testing.T) {
	suite.Run(t, new(LogoutServiceTestSuite))
}

type LogoutServiceTestSuite struct {
	suite.Suite
	fakeAuthService *mockery.AuthServiceInterface
	authConfig      config.AuthConfig
	service         auth.LogoutService
	context         utils.CommonContext
}

func (suite *LogoutServiceTestSuite) SetupTest() {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Set(constants.AccessTokenContext, "token")
	suite.context = utils.NewCommonContext(c, nil, nil)
	suite.fakeAuthService = mockery.NewAuthServiceInterface(suite.T())
	suite.authConfig = config.AuthConfig{}
	suite.service = auth.NewLogoutService(suite.context, suite.fakeAuthService, suite.authConfig)
}

func (suite *LogoutServiceTestSuite) TestLogout() {
	suite.fakeAuthService.EXPECT().InvalidateToken(mock.Anything, "token").Return(nil)

	err := suite.service.ProcessingLogout()
	suite.Nil(err)
}
