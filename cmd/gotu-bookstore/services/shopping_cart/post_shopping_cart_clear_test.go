package shopping_cart_test

import (
	"gotu-bookstore/cmd/gotu-bookstore/constants"
	"gotu-bookstore/cmd/gotu-bookstore/models"
	"gotu-bookstore/cmd/gotu-bookstore/services/shopping_cart"
	mockery "gotu-bookstore/mocks/cmd/gotu-bookstore/services/shopping_cart"
	"gotu-bookstore/pkg/auth/dto"
	"gotu-bookstore/pkg/resfmt/base_error"
	"gotu-bookstore/pkg/utils"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

func TestPostShoppingCartClearServiceTestSuite(t *testing.T) {
	suite.Run(t, new(PostShoppingCartClearServiceTestSuite))
}

type PostShoppingCartClearServiceTestSuite struct {
	suite.Suite
	fakeShoppingCartRepo *mockery.ShoppingCartRepoInterface
	fakeUserRepo         *mockery.UserRepoInterface
	service              shopping_cart.PostShoppingCartClearService
	context              utils.CommonContext
	sessionDTO           dto.SessionDTO
	user                 models.Users
}

func (suite *PostShoppingCartClearServiceTestSuite) SetupTest() {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
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
	suite.fakeShoppingCartRepo = mockery.NewShoppingCartRepoInterface(suite.T())
	suite.fakeUserRepo = mockery.NewUserRepoInterface(suite.T())
	suite.service = shopping_cart.NewPostShoppingCartClearService(suite.context, suite.fakeShoppingCartRepo, suite.fakeUserRepo)
}

func (suite *PostShoppingCartClearServiceTestSuite) TestClearShoppingCart() {
	suite.fakeUserRepo.EXPECT().GetById(suite.sessionDTO.Id).Return(&suite.user, nil)
	suite.fakeShoppingCartRepo.EXPECT().ClearShoppingCart(suite.sessionDTO.Id).Return(nil)
	response, err := suite.service.ProcessingPostShoppingCartClear()
	suite.Nil(err)
	suite.Nil(response)
}

func (suite *PostShoppingCartClearServiceTestSuite) TestClearShoppingCart_Error() {
	err := base_error.New("error")
	suite.fakeUserRepo.EXPECT().GetById(suite.sessionDTO.Id).Return(&suite.user, nil)
	suite.fakeShoppingCartRepo.EXPECT().ClearShoppingCart(suite.sessionDTO.Id).Return(err)
	response, err := suite.service.ProcessingPostShoppingCartClear()
	suite.NotNil(err)
	suite.Nil(response)
}
