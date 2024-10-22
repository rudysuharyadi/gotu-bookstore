package shopping_cart_test

import (
	"gotu-bookstore/cmd/gotu-bookstore/constants"
	"gotu-bookstore/cmd/gotu-bookstore/models"
	"gotu-bookstore/cmd/gotu-bookstore/services/shopping_cart"
	mockery "gotu-bookstore/mocks/cmd/gotu-bookstore/services/shopping_cart"
	"gotu-bookstore/pkg/auth/dto"
	"gotu-bookstore/pkg/utils"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

func TestGetShoppingCartServiceTestSuite(t *testing.T) {
	suite.Run(t, new(GetShoppingCartServiceTestSuite))
}

type GetShoppingCartServiceTestSuite struct {
	suite.Suite
	fakeShoppingCartRepo *mockery.ShoppingCartRepoInterface
	fakeUserRepo         *mockery.UserRepoInterface
	service              shopping_cart.GetShoppingCartService
	context              utils.CommonContext
	sessionDTO           dto.SessionDTO
	user                 models.Users
}

func (suite *GetShoppingCartServiceTestSuite) SetupTest() {
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
	suite.service = shopping_cart.NewGetShoppingCartService(suite.context, suite.fakeShoppingCartRepo, suite.fakeUserRepo)
}

func (suite *GetShoppingCartServiceTestSuite) TestGetShoppingCart() {
	cartItems := GenerateFakerShoppingCart(3)
	suite.fakeUserRepo.EXPECT().GetById(suite.sessionDTO.Id).Return(&suite.user, nil)
	suite.fakeShoppingCartRepo.EXPECT().GetByUserId(suite.sessionDTO.Id).Return(cartItems, nil)
	response, err := suite.service.ProcessingGetShoppingCart()
	suite.Nil(err)
	suite.NotNil(response)
	suite.Equal(len(cartItems), len(response.Items))
	grandTotal := 0.0
	for _, cartItem := range cartItems {
		grandTotal += cartItem.Book.Price
	}
	suite.Equal(utils.FloatToString(grandTotal), response.GrandTotal)
}
