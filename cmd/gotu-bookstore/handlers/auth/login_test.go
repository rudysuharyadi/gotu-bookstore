package auth_test

import (
	"fmt"
	"gotu-bookstore/cmd/gotu-bookstore/handlers/auth"
	"gotu-bookstore/pkg/utils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	contracts "gotu-bookstore/cmd/gotu-bookstore/contracts/auth"
	mockery "gotu-bookstore/mocks/cmd/gotu-bookstore/handlers/auth"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestLoginHandlerSuite(t *testing.T) {
	suite.Run(t, new(LoginHandlerTestSuite))
}

type LoginHandlerTestSuite struct {
	suite.Suite
	fakeLoginService *mockery.LoginServiceInterface
	handler          auth.LoginHandler
	context          utils.CommonContext
	engine           *gin.Engine
	recorder         *httptest.ResponseRecorder
}

func (suite *LoginHandlerTestSuite) SetupTest() {
	suite.recorder = httptest.NewRecorder()
	_, engine := gin.CreateTestContext(suite.recorder)
	suite.engine = engine
	suite.fakeLoginService = mockery.NewLoginServiceInterface(suite.T())
	suite.engine.POST("/account-service/v1/login", func(c *gin.Context) {
		suite.context = utils.NewCommonContext(c, nil, nil)
		suite.handler = auth.NewLoginHandler(suite.context, suite.fakeLoginService)
		suite.handler.ProcessingLogin()
	})
}

func (suite *LoginHandlerTestSuite) TestLogin() {
	email := "rudy.suharyadi@gmail.com"
	password := "pass"

	request, _ := http.NewRequest(
		"POST",
		"/account-service/v1/login",
		strings.NewReader(fmt.Sprintf(`{
			"email": "%s",
			"password": "%s"
		}`, email, password)),
	)
	appendHeaders(request)

	var capturedLoginRequest contracts.LoginRequest
	suite.fakeLoginService.EXPECT().
		ProcessingLogin(mock.MatchedBy(func(arg contracts.LoginRequest) bool {
			capturedLoginRequest = arg
			return true
		})).
		Return(nil, nil)

	suite.engine.ServeHTTP(suite.recorder, request)
	suite.Equal(http.StatusOK, suite.recorder.Code)
	suite.Equal(email, capturedLoginRequest.Email)
	suite.Equal(password, capturedLoginRequest.Password)
}

func (suite *LoginHandlerTestSuite) TestLogin_ErrorBadRequest() {
	request, _ := http.NewRequest(
		"POST",
		"/account-service/v1/login",
		strings.NewReader(``),
	)
	appendHeaders(request)

	suite.engine.ServeHTTP(suite.recorder, request)
	suite.Equal(http.StatusBadRequest, suite.recorder.Code)
}

func appendHeaders(req *http.Request) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	for name, value := range headers {
		req.Header.Add(name, value)
	}
}
