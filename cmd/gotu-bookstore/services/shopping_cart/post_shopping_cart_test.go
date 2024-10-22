package shopping_cart_test

import (
	"gotu-bookstore/cmd/gotu-bookstore/constants"
	contracts "gotu-bookstore/cmd/gotu-bookstore/contracts/shopping_cart"
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

func TestPostShoppingCartServiceTestSuite(t *testing.T) {
	suite.Run(t, new(PostShoppingCartServiceTestSuite))
}

type PostShoppingCartServiceTestSuite struct {
	suite.Suite
	fakeShoppingCartRepo *mockery.ShoppingCartRepoInterface
	fakeUserRepo         *mockery.UserRepoInterface
	fakeBookRepo         *mockery.BookRepoInterface
	service              shopping_cart.PostShoppingCartService
	context              utils.CommonContext
	sessionDTO           dto.SessionDTO
	user                 models.Users
}

func (suite *PostShoppingCartServiceTestSuite) SetupTest() {
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
	suite.fakeBookRepo = mockery.NewBookRepoInterface(suite.T())
	suite.service = shopping_cart.NewPostShoppingCartService(
		suite.context, suite.fakeShoppingCartRepo, suite.fakeUserRepo, suite.fakeBookRepo)
}

func (suite *PostShoppingCartServiceTestSuite) TestAddItemToShoppingCart() {
	book := GenerateFakeBook(1)[0]
	request := contracts.PostShoppingCartRequest{
		BookId:   book.Id.String(),
		Quantity: 2,
	}

	cartItems := GenerateFakerShoppingCart(1)
	var capturedItems models.ShoppingCarts

	suite.fakeUserRepo.EXPECT().GetById(suite.sessionDTO.Id).Return(&suite.user, nil)
	suite.fakeBookRepo.EXPECT().GetById(book.Id.String()).Return(&book, nil)
	suite.fakeShoppingCartRepo.EXPECT().GetItemByBookId(book.Id.String(), suite.sessionDTO.Id).
		Return(nil, nil)
	suite.fakeShoppingCartRepo.EXPECT().
		AddItemToCart(mock.MatchedBy(func(arg models.ShoppingCarts) bool {
			capturedItems = arg
			return true
		})).
		Return(nil, nil)
	suite.fakeShoppingCartRepo.EXPECT().GetByUserId(suite.sessionDTO.Id).Return(cartItems, nil)

	response, err := suite.service.ProcessingPostShoppingCart(request)

	suite.Nil(err)
	suite.NotNil(response)
	suite.Equal(len(cartItems), len(response.Items))
	grandTotal := 0.0
	for _, cartItem := range cartItems {
		grandTotal += cartItem.Book.Price
	}
	suite.Equal(utils.FloatToString(grandTotal), response.GrandTotal)
	suite.Equal(int64(request.Quantity), capturedItems.Quantity)
	suite.Equal(request.BookId, capturedItems.BookId.String())
}

func (suite *PostShoppingCartServiceTestSuite) TestEditItemOnShoppingCart() {
	book := GenerateFakeBook(1)[0]
	request := contracts.PostShoppingCartRequest{
		BookId:   book.Id.String(),
		Quantity: 2,
	}

	cartItems := GenerateFakerShoppingCart(1)
	cartItem := cartItems[0]

	suite.fakeUserRepo.EXPECT().GetById(suite.sessionDTO.Id).Return(&suite.user, nil)
	suite.fakeBookRepo.EXPECT().GetById(book.Id.String()).Return(&book, nil)
	suite.fakeShoppingCartRepo.EXPECT().GetItemByBookId(book.Id.String(), suite.sessionDTO.Id).
		Return(&cartItem, nil)
	suite.fakeShoppingCartRepo.EXPECT().
		UpdateByBookId(request.BookId, int64(request.Quantity), suite.sessionDTO.Id).
		Return(nil)
	suite.fakeShoppingCartRepo.EXPECT().GetByUserId(suite.sessionDTO.Id).Return(cartItems, nil)

	response, err := suite.service.ProcessingPostShoppingCart(request)

	suite.Nil(err)
	suite.NotNil(response)
	suite.Equal(len(cartItems), len(response.Items))
	grandTotal := 0.0
	for _, cartItem := range cartItems {
		grandTotal += cartItem.Book.Price
	}
	suite.Equal(utils.FloatToString(grandTotal), response.GrandTotal)
}

func (suite *PostShoppingCartServiceTestSuite) TestDeleteItemOnShoppingCart() {
	book := GenerateFakeBook(1)[0]
	request := contracts.PostShoppingCartRequest{
		BookId:   book.Id.String(),
		Quantity: 0,
	}

	cartItems := GenerateFakerShoppingCart(1)
	cartItem := cartItems[0]

	suite.fakeUserRepo.EXPECT().GetById(suite.sessionDTO.Id).Return(&suite.user, nil)
	suite.fakeBookRepo.EXPECT().GetById(book.Id.String()).Return(&book, nil)
	suite.fakeShoppingCartRepo.EXPECT().GetItemByBookId(book.Id.String(), suite.sessionDTO.Id).
		Return(&cartItem, nil)
	suite.fakeShoppingCartRepo.EXPECT().
		DeleteByBookId(request.BookId, suite.sessionDTO.Id).
		Return(nil)
	suite.fakeShoppingCartRepo.EXPECT().GetByUserId(suite.sessionDTO.Id).Return(cartItems, nil)

	response, err := suite.service.ProcessingPostShoppingCart(request)

	suite.Nil(err)
	suite.NotNil(response)
	suite.Equal(len(cartItems), len(response.Items))
	grandTotal := 0.0
	for _, cartItem := range cartItems {
		grandTotal += cartItem.Book.Price
	}
	suite.Equal(utils.FloatToString(grandTotal), response.GrandTotal)
}
