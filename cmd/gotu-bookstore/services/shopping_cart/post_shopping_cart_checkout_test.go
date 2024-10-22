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
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestPostShoppingCartCheckoutServiceTestSuite(t *testing.T) {
	suite.Run(t, new(PostShoppingCartCheckoutServiceTestSuite))
}

type PostShoppingCartCheckoutServiceTestSuite struct {
	suite.Suite
	fakeShoppingCartRepo *mockery.ShoppingCartRepoInterface
	fakeUserRepo         *mockery.UserRepoInterface
	fakeTransactionRepo  *mockery.TransactionRepoInterface
	service              shopping_cart.PostShoppingCartCheckoutService
	context              utils.CommonContext
	sessionDTO           dto.SessionDTO
	user                 models.Users
}

func (suite *PostShoppingCartCheckoutServiceTestSuite) SetupTest() {
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
	suite.fakeTransactionRepo = mockery.NewTransactionRepoInterface(suite.T())
	suite.service = shopping_cart.NewPostShoppingCartCheckoutService(
		suite.context, suite.fakeShoppingCartRepo, suite.fakeUserRepo, suite.fakeTransactionRepo)
}

func (suite *PostShoppingCartCheckoutServiceTestSuite) TestCheckoutShoppingCart() {
	var capturedTransaction models.Transactions
	cartItems := GenerateFakerShoppingCart(3)

	suite.fakeUserRepo.EXPECT().GetById(suite.sessionDTO.Id).Return(&suite.user, nil)
	suite.fakeShoppingCartRepo.EXPECT().GetByUserId(suite.sessionDTO.Id).Return(cartItems, nil)
	suite.fakeTransactionRepo.EXPECT().GenerateInvoiceCounter().Return(1, nil)
	suite.fakeTransactionRepo.EXPECT().
		Create(mock.MatchedBy(func(arg models.Transactions) bool {
			capturedTransaction = arg
			return true
		})).
		Return(nil, nil)
	suite.fakeShoppingCartRepo.EXPECT().ClearShoppingCart(suite.sessionDTO.Id).Return(nil)

	response, err := suite.service.ProcessingPostShoppingCartCheckout()

	suite.Nil(err)
	suite.Nil(response)
	suite.Equal("k", capturedTransaction.InvoiceNumber)
	suite.Equal(suite.sessionDTO.Id, capturedTransaction.UserId.String())
	suite.Equal(string(constants.TransactionStatusConfirmed), capturedTransaction.Status)
	suite.Equal(len(cartItems), len(capturedTransaction.TransactionItems))

	grandTotal := 0.0
	for _, cartItem := range cartItems {
		grandTotal += cartItem.Book.Price
	}
	suite.Equal(grandTotal, capturedTransaction.GrandTotal)
}
