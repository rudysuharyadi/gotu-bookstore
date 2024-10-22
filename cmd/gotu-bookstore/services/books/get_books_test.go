package books_test

import (
	"errors"
	contracts "gotu-bookstore/cmd/gotu-bookstore/contracts/books"
	"gotu-bookstore/cmd/gotu-bookstore/contracts/pagination"
	"gotu-bookstore/cmd/gotu-bookstore/models"
	"gotu-bookstore/cmd/gotu-bookstore/services/books"
	mockery "gotu-bookstore/mocks/cmd/gotu-bookstore/services/books"
	"gotu-bookstore/pkg/resfmt/base_error"
	"gotu-bookstore/pkg/utils"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

func TestGetBooksServiceSuite(t *testing.T) {
	suite.Run(t, new(GetBooksServiceTestSuite))
}

type GetBooksServiceTestSuite struct {
	suite.Suite
	fakeBookRepo *mockery.BooksRepoInterface
	service      books.GetBooksService
	context      utils.CommonContext
}

func (suite *GetBooksServiceTestSuite) SetupTest() {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	suite.context = utils.NewCommonContext(c, nil, nil)
	suite.fakeBookRepo = mockery.NewBooksRepoInterface(suite.T())
	suite.service = books.NewGetBooksService(suite.context, suite.fakeBookRepo)
}

func (suite *GetBooksServiceTestSuite) TestProcessingGetBooks() {
	limit := 3
	page := 0
	sortBy := ""
	desc := false
	keyword := ""
	books := CreateTestBooks(3)
	size := int64(len(books))
	var err error

	suite.fakeBookRepo.EXPECT().GetAll(limit, page, sortBy, desc, keyword).
		Return(books, size, err)

	request := contracts.GetBooksRequest{
		Pagination: pagination.Pagination{
			Page:    page,
			Limit:   limit,
			SortBy:  sortBy,
			Desc:    desc,
			Keyword: keyword,
		},
	}
	result, metadata, err := suite.service.ProcessingGetBooks(request)

	suite.NotNil(result)
	suite.NotNil(metadata)
	suite.Nil(err)
	suite.Equal(3, len(*result))
}

func (suite *GetBooksServiceTestSuite) TestProcessingGetBooks_ErrorWhenGetBookData() {
	limit := 3
	page := 0
	sortBy := ""
	desc := false
	keyword := ""
	books := []models.Books{}
	size := int64(0)
	err := base_error.New("Error when get book data")

	suite.fakeBookRepo.EXPECT().GetAll(limit, page, sortBy, desc, keyword).
		Return(books, size, err)

	request := contracts.GetBooksRequest{
		Pagination: pagination.Pagination{
			Page:    page,
			Limit:   limit,
			SortBy:  sortBy,
			Desc:    desc,
			Keyword: keyword,
		},
	}
	result, metadata, err := suite.service.ProcessingGetBooks(request)

	suite.Nil(result)
	suite.Nil(metadata)
	suite.NotNil(err)
	suite.True(errors.As(err, &base_error.BaseError{}))
}
