package books_test

import (
	contracts "gotu-bookstore/cmd/gotu-bookstore/contracts/books"
	"gotu-bookstore/cmd/gotu-bookstore/services/books"
	mockery "gotu-bookstore/mocks/cmd/gotu-bookstore/services/books"
	"gotu-bookstore/pkg/resfmt/base_error"
	"gotu-bookstore/pkg/utils"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

func TestGetBookDetailsServiceSuite(t *testing.T) {
	suite.Run(t, new(GetBooksServiceTestSuite))
}

type GetBookDetailsServiceTestSuite struct {
	suite.Suite
	fakeBookRepo *mockery.BooksRepoInterface
	service      books.GetBookDetailsService
	context      utils.CommonContext
}

func (suite *GetBookDetailsServiceTestSuite) SetupTest() {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	suite.context = utils.NewCommonContext(c, nil, nil)
	suite.fakeBookRepo = mockery.NewBooksRepoInterface(suite.T())
	suite.service = books.NewGetBookDetailsService(suite.context, suite.fakeBookRepo)
}

func (suite *GetBookDetailsServiceTestSuite) TestGetBookDetails() {
	book := CreateTestBooks(1)[0]
	request := contracts.GetBookDetailsRequest{
		BookId: book.Id.String(),
	}

	suite.fakeBookRepo.EXPECT().GetById(request.BookId).
		Return(&book, nil)

	response, err := suite.service.ProcessingGetBookDetails(request)
	suite.Nil(err)
	suite.NotNil(response)
	suite.Equal(book.Id.String(), response.BooksDTO.Id)
	suite.Equal(utils.FloatToString(book.Price), response.BooksDTO.Price)
	suite.Equal(book.Title, response.BooksDTO.Title)
}

func (suite *GetBookDetailsServiceTestSuite) TestGetBookDetails_Error() {
	request := contracts.GetBookDetailsRequest{
		BookId: uuid.New().String(),
	}
	err := base_error.New("error")

	suite.fakeBookRepo.EXPECT().GetById(request.BookId).
		Return(nil, err)

	response, err := suite.service.ProcessingGetBookDetails(request)
	suite.NotNil(err)
	suite.Nil(response)
}
