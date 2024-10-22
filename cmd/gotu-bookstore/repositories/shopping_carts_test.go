package repositories_test

import (
	"gotu-bookstore/cmd/gotu-bookstore/models"
	"gotu-bookstore/cmd/gotu-bookstore/repositories"
	"gotu-bookstore/pkg/database"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

func TestShoppingCartsRepositorySuite(t *testing.T) {
	suite.Run(t, new(ShoppingCartsRepositoryTestSuite))
}

type ShoppingCartsRepositoryTestSuite struct {
	sqlmock    sqlmock.Sqlmock
	repository repositories.ShoppingCartsRepository
	suite.Suite
}

func (suite *ShoppingCartsRepositoryTestSuite) SetupTest() {
	db, mock, err := database.NewMockDB()
	suite.Nil(err)
	suite.sqlmock = mock
	suite.repository = repositories.NewShoppingCartsRepository(db)
}

func (suite *ShoppingCartsRepositoryTestSuite) TestGetByUserId() {
	id := uuid.New()
	userId := uuid.New()
	bookId := uuid.New()
	quantity := 1
	title := "title"
	author := "author"
	price := 10.59

	selectQuery := regexp.QuoteMeta(`SELECT * FROM "shopping_carts" WHERE user_id = $1`)
	suite.sqlmock.ExpectQuery(selectQuery).
		WithArgs(userId).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_id", "book_id", "quantity"}).
				AddRow(id, userId, bookId, quantity),
		)

	selectQuery = regexp.QuoteMeta(`SELECT * FROM "books" WHERE "books"."id" = $1`)
	suite.sqlmock.ExpectQuery(selectQuery).
		WithArgs(bookId).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "author", "price"}).
				AddRow(bookId, title, author, price),
		)

	result, err := suite.repository.GetByUserId(userId.String())
	suite.Nil(err)
	suite.NotNil(result)
	suite.Equal(id, result[0].Id)
	suite.Equal(userId, result[0].UserId)
	suite.Equal(bookId, result[0].BookId)
	suite.Equal(int64(quantity), result[0].Quantity)

	err = suite.sqlmock.ExpectationsWereMet()
	suite.NoError(err)
}

func (suite *ShoppingCartsRepositoryTestSuite) TestItemByBookId() {
	id := uuid.New()
	userId := uuid.New()
	bookId := uuid.New()
	quantity := 1

	selectQuery := regexp.QuoteMeta(`SELECT * FROM "shopping_carts" WHERE (book_id = $1 AND user_id = $2)`)
	suite.sqlmock.ExpectQuery(selectQuery).
		WithArgs(bookId, userId, 1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_id", "book_id", "quantity"}).
				AddRow(id, userId, bookId, quantity),
		)

	result, err := suite.repository.GetItemByBookId(bookId.String(), userId.String())
	suite.Nil(err)
	suite.NotNil(result)
	suite.Equal(id, result.Id)
	suite.Equal(userId, result.UserId)
	suite.Equal(bookId, result.BookId)
	suite.Equal(int64(quantity), result.Quantity)

	err = suite.sqlmock.ExpectationsWereMet()
	suite.NoError(err)
}

func (suite *ShoppingCartsRepositoryTestSuite) TestAddItemToCart() {
	input := models.ShoppingCarts{
		Id:        uuid.New(),
		UserId:    uuid.New(),
		BookId:    uuid.New(),
		Quantity:  1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	suite.sqlmock.ExpectBegin()
	suite.sqlmock.ExpectExec(`INSERT INTO "shopping_carts"`).
		WithArgs(
			input.Id,
			input.UserId,
			input.BookId,
			input.Quantity,
			input.CreatedAt,
			input.UpdatedAt,
			input.DeletedAt,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.sqlmock.ExpectCommit()
	result, err := suite.repository.AddItemToCart(input)
	suite.Nil(err)
	suite.NotNil(result)

	err = suite.sqlmock.ExpectationsWereMet()
	suite.NoError(err)
}

func (suite *ShoppingCartsRepositoryTestSuite) TestUpdateByBookId() {
	userId := uuid.New()
	suite.sqlmock.ExpectBegin()
	suite.sqlmock.ExpectExec(`UPDATE "shopping_carts"`).
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.sqlmock.ExpectCommit()
	err := suite.repository.ClearShoppingCart(userId.String())
	suite.Nil(err)

	err = suite.sqlmock.ExpectationsWereMet()
	suite.NoError(err)
}

func (suite *ShoppingCartsRepositoryTestSuite) TestDeleteByBookId() {
	userId := uuid.New()
	suite.sqlmock.ExpectBegin()
	suite.sqlmock.ExpectExec(`UPDATE "shopping_carts"`).
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.sqlmock.ExpectCommit()
	err := suite.repository.ClearShoppingCart(userId.String())
	suite.Nil(err)

	err = suite.sqlmock.ExpectationsWereMet()
	suite.NoError(err)
}

func (suite *ShoppingCartsRepositoryTestSuite) TestClearShoppingCart() {
	userId := uuid.New()
	suite.sqlmock.ExpectBegin()
	suite.sqlmock.ExpectExec(`UPDATE "shopping_carts"`).
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.sqlmock.ExpectCommit()
	err := suite.repository.ClearShoppingCart(userId.String())
	suite.Nil(err)

	err = suite.sqlmock.ExpectationsWereMet()
	suite.NoError(err)
}
