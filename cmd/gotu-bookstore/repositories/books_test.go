package repositories_test

import (
	"gotu-bookstore/cmd/gotu-bookstore/repositories"
	"gotu-bookstore/pkg/database"
	"gotu-bookstore/pkg/resfmt/base_error"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

func TestBooksRepositorySuite(t *testing.T) {
	suite.Run(t, new(BooksRepositoryTestSuite))
}

type BooksRepositoryTestSuite struct {
	sqlmock    sqlmock.Sqlmock
	repository repositories.BooksRepository
	suite.Suite
}

func (suite *BooksRepositoryTestSuite) SetupTest() {
	db, mock, err := database.NewMockDB()
	suite.Nil(err)
	suite.sqlmock = mock
	suite.repository = repositories.NewBooksRepository(db)
}

func (suite *BooksRepositoryTestSuite) TestGetAllBooks() {
	limit := 3
	page := 0
	sortBy := ""
	desc := false
	keyword := ""

	countQuery := regexp.QuoteMeta(`SELECT count(*) FROM "books"`)
	suite.sqlmock.ExpectQuery(countQuery).
		WillReturnRows(suite.sqlmock.NewRows([]string{"count"}).AddRow(2))

	selectQuery := regexp.QuoteMeta(`SELECT * FROM "books" WHERE "books"."deleted_at" IS NULL LIMIT $1`)
	suite.sqlmock.ExpectQuery(selectQuery).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "author", "isbn", "description"}).
				AddRow(uuid.New(), "Test Book 1", "Author 1", "1234567890", "Description 1").
				AddRow(uuid.New(), "Test Book 2", "Author 2", "0987654321", "Description 2"),
		)

	result, size, err := suite.repository.GetAll(limit, page, sortBy, desc, keyword)
	suite.Nil(err)
	suite.Equal(size, int64(2))
	suite.Len(result, 2)
	suite.Equal("Test Book 1", result[0].Title)

	err = suite.sqlmock.ExpectationsWereMet()
	suite.NoError(err)
}

func (suite *BooksRepositoryTestSuite) TestGetAllBooksWithSortBy() {
	limit := 3
	page := 0
	sortBy := "title"
	desc := false
	keyword := ""

	countQuery := regexp.QuoteMeta(`SELECT count(*) FROM "books"`)
	suite.sqlmock.ExpectQuery(countQuery).
		WillReturnRows(suite.sqlmock.NewRows([]string{"count"}).AddRow(2))

	selectQuery := regexp.QuoteMeta(`SELECT * FROM "books" WHERE "books"."deleted_at" IS NULL ORDER BY "title" LIMIT $1`)
	suite.sqlmock.ExpectQuery(selectQuery).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "author", "isbn", "description"}).
				AddRow(uuid.New(), "Test Book 1", "Author 1", "1234567890", "Description 1").
				AddRow(uuid.New(), "Test Book 2", "Author 2", "0987654321", "Description 2"),
		)

	result, size, err := suite.repository.GetAll(limit, page, sortBy, desc, keyword)
	suite.Nil(err)
	suite.Equal(size, int64(2))
	suite.Len(result, 2)
	suite.Equal("Test Book 1", result[0].Title)

	err = suite.sqlmock.ExpectationsWereMet()
	suite.NoError(err)
}

func (suite *BooksRepositoryTestSuite) TestGetAllBooksWithSortByDesc() {
	limit := 3
	page := 0
	sortBy := "title"
	desc := true
	keyword := ""

	countQuery := regexp.QuoteMeta(`SELECT count(*) FROM "books"`)
	suite.sqlmock.ExpectQuery(countQuery).
		WillReturnRows(suite.sqlmock.NewRows([]string{"count"}).AddRow(2))

	selectQuery := regexp.QuoteMeta(`SELECT * FROM "books" WHERE "books"."deleted_at" IS NULL ORDER BY "title" DESC LIMIT $1`)
	suite.sqlmock.ExpectQuery(selectQuery).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "author", "isbn", "description"}).
				AddRow(uuid.New(), "Test Book 1", "Author 1", "1234567890", "Description 1").
				AddRow(uuid.New(), "Test Book 2", "Author 2", "0987654321", "Description 2"),
		)

	result, size, err := suite.repository.GetAll(limit, page, sortBy, desc, keyword)
	suite.Nil(err)
	suite.Equal(size, int64(2))
	suite.Len(result, 2)
	suite.Equal("Test Book 1", result[0].Title)

	err = suite.sqlmock.ExpectationsWereMet()
	suite.NoError(err)
}

func (suite *BooksRepositoryTestSuite) TestGetAllBooksWithSortBy_InvalidColumn() {
	limit := 3
	page := 0
	sortBy := "invalid"
	desc := true
	keyword := ""

	countQuery := regexp.QuoteMeta(`SELECT count(*) FROM "books"`)
	suite.sqlmock.ExpectQuery(countQuery).
		WillReturnRows(suite.sqlmock.NewRows([]string{"count"}).AddRow(2))

	result, size, err := suite.repository.GetAll(limit, page, sortBy, desc, keyword)
	suite.NotNil(err)
	suite.Empty(result)
	suite.Equal(int64(0), size)
}

func (suite *BooksRepositoryTestSuite) TestGetAllBooksWithKeyword() {
	limit := 3
	page := 0
	sortBy := ""
	desc := false
	keyword := "Author"

	countQuery := regexp.QuoteMeta(`SELECT count(*) FROM "books"`)
	suite.sqlmock.ExpectQuery(countQuery).
		WillReturnRows(suite.sqlmock.NewRows([]string{"count"}).AddRow(2))

	selectQuery := regexp.QuoteMeta(`SELECT * FROM "books" WHERE 
		(LOWER(title) LIKE $1 OR 
		LOWER(author) LIKE $2 OR 
		LOWER(description) LIKE $3 OR 
		LOWER(isbn) LIKE $4) AND "books"."deleted_at" IS NULL LIMIT $5`)
	suite.sqlmock.ExpectQuery(selectQuery).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "author", "isbn", "description"}).
				AddRow(uuid.New(), "Test Book 1", "Author 1", "1234567890", "Description 1").
				AddRow(uuid.New(), "Test Book 2", "Author 2", "0987654321", "Description 2"),
		)

	result, size, err := suite.repository.GetAll(limit, page, sortBy, desc, keyword)
	suite.Nil(err)
	suite.Equal(size, int64(2))
	suite.Len(result, 2)
	suite.Equal("Test Book 1", result[0].Title)

	err = suite.sqlmock.ExpectationsWereMet()
	suite.NoError(err)
}

func (suite *BooksRepositoryTestSuite) TestGetById() {
	id := uuid.New()
	title := "Test Book 1"
	author := "Author 1"
	isbn := "1234567890"
	description := "Description 1"

	selectQuery := regexp.QuoteMeta(`SELECT * FROM "books" WHERE id = $1`)
	suite.sqlmock.ExpectQuery(selectQuery).
		WithArgs(id, 1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "author", "isbn", "description"}).
				AddRow(id, title, author, isbn, description),
		)

	result, err := suite.repository.GetById(id.String())
	suite.Nil(err)
	suite.NotNil(result)
	suite.Equal(author, result.Author)
	suite.Equal(title, result.Title)
	suite.Equal(description, result.Description)
	suite.Equal(isbn, result.Isbn)
	suite.Equal(id, result.Id)

	err = suite.sqlmock.ExpectationsWereMet()
	suite.NoError(err)
}

func (suite *BooksRepositoryTestSuite) TestGetById_Error() {
	id := uuid.New()
	err := base_error.New("test")

	selectQuery := regexp.QuoteMeta(`SELECT * FROM "books" WHERE id = $1`)
	suite.sqlmock.ExpectQuery(selectQuery).
		WithArgs(id, 1).
		WillReturnError(err)

	result, err := suite.repository.GetById(id.String())
	suite.NotNil(err)
	suite.Empty(result)

	err = suite.sqlmock.ExpectationsWereMet()
	suite.NoError(err)
}
