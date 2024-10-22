package repositories_test

import (
	"gotu-bookstore/cmd/gotu-bookstore/constants"
	"gotu-bookstore/cmd/gotu-bookstore/models"
	"gotu-bookstore/cmd/gotu-bookstore/repositories"
	"gotu-bookstore/pkg/database"
	"gotu-bookstore/pkg/resfmt/base_error"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

func TestTransactionsRepositorySuite(t *testing.T) {
	suite.Run(t, new(TransactionsRepositoryTestSuite))
}

type TransactionsRepositoryTestSuite struct {
	sqlmock    sqlmock.Sqlmock
	repository repositories.TransactionsRepository
	suite.Suite
}

func (suite *TransactionsRepositoryTestSuite) SetupTest() {
	db, mock, err := database.NewMockDB()
	suite.Nil(err)
	suite.sqlmock = mock
	suite.repository = repositories.NewTransactionsRepository(db)
}

func (suite *TransactionsRepositoryTestSuite) TestGetAll() {
	limit := 3
	page := 0
	sortBy := ""
	desc := false

	id := uuid.New()
	userId := uuid.New()
	grandTotal := 10.59
	status := string(constants.TransactionStatusConfirmed)
	invoiceNumber := "abcd"

	itemId := uuid.New()
	bookId := uuid.New()
	quantity := int64(1)
	price := 10.59

	// Expect count transactions
	countQuery := regexp.QuoteMeta(`SELECT count(*) FROM "transactions" WHERE user_id = $1`)
	suite.sqlmock.ExpectQuery(countQuery).
		WithArgs(userId).
		WillReturnRows(suite.sqlmock.NewRows([]string{"count"}).AddRow(1))

	// Expect select transactions table
	selectQuery := regexp.QuoteMeta(`SELECT * FROM "transactions" WHERE user_id = $1 AND "transactions"."deleted_at" IS NULL LIMIT $2`)
	suite.sqlmock.ExpectQuery(selectQuery).
		WithArgs(userId, limit).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_id", "grand_total", "status", "invoice_number"}).
				AddRow(id, userId, grandTotal, status, invoiceNumber),
		)

	// Expect select transaction_items table with Preload() functions.
	selectQuery = regexp.QuoteMeta(`SELECT * FROM "transaction_items" WHERE "transaction_items"."transaction_id" = $1`)
	suite.sqlmock.ExpectQuery(selectQuery).
		WithArgs(id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "transaction_id", "book_id", "quantity", "price"}).
				AddRow(itemId, id, bookId, quantity, price),
		)

	result, size, err := suite.repository.GetAll(userId.String(), limit, page, sortBy, desc)
	suite.Nil(err)
	suite.NotNil(result)
	suite.Equal(int64(1), size)
	suite.Equal(userId, result[0].UserId)
	suite.Equal(id, result[0].Id)
	suite.Equal(grandTotal, result[0].GrandTotal)
	suite.Equal(status, result[0].Status)
	suite.Equal(invoiceNumber, result[0].InvoiceNumber)
	suite.Equal(quantity, result[0].TransactionItems[0].Quantity)
	suite.Equal(price, result[0].TransactionItems[0].Price)
	suite.Equal(bookId, result[0].TransactionItems[0].BookId)
	suite.Equal(itemId, result[0].TransactionItems[0].Id)

	err = suite.sqlmock.ExpectationsWereMet()
	suite.NoError(err)
}

func (suite *TransactionsRepositoryTestSuite) TestGetAllWithSortBy() {
	limit := 3
	page := 0
	sortBy := "created_at"
	desc := false

	id := uuid.New()
	userId := uuid.New()
	grandTotal := 10.59
	status := string(constants.TransactionStatusConfirmed)
	invoiceNumber := "abcd"

	itemId := uuid.New()
	bookId := uuid.New()
	quantity := int64(1)
	price := 10.59

	// Expect count transactions
	countQuery := regexp.QuoteMeta(`SELECT count(*) FROM "transactions" WHERE user_id = $1`)
	suite.sqlmock.ExpectQuery(countQuery).
		WithArgs(userId).
		WillReturnRows(suite.sqlmock.NewRows([]string{"count"}).AddRow(1))

	// Expect select transactions table
	selectQuery := regexp.QuoteMeta(`SELECT * FROM "transactions" WHERE user_id = $1 AND "transactions"."deleted_at" IS NULL ORDER BY "created_at" LIMIT $2`)
	suite.sqlmock.ExpectQuery(selectQuery).
		WithArgs(userId, limit).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_id", "grand_total", "status", "invoice_number"}).
				AddRow(id, userId, grandTotal, status, invoiceNumber),
		)

	// Expect select transaction_items table with Preload() functions.
	selectQuery = regexp.QuoteMeta(`SELECT * FROM "transaction_items" WHERE "transaction_items"."transaction_id" = $1`)
	suite.sqlmock.ExpectQuery(selectQuery).
		WithArgs(id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "transaction_id", "book_id", "quantity", "price"}).
				AddRow(itemId, id, bookId, quantity, price),
		)

	result, size, err := suite.repository.GetAll(userId.String(), limit, page, sortBy, desc)
	suite.Nil(err)
	suite.NotNil(result)
	suite.Equal(int64(1), size)
	suite.Equal(userId, result[0].UserId)
	suite.Equal(id, result[0].Id)
	suite.Equal(grandTotal, result[0].GrandTotal)
	suite.Equal(status, result[0].Status)
	suite.Equal(invoiceNumber, result[0].InvoiceNumber)
	suite.Equal(quantity, result[0].TransactionItems[0].Quantity)
	suite.Equal(price, result[0].TransactionItems[0].Price)
	suite.Equal(bookId, result[0].TransactionItems[0].BookId)
	suite.Equal(itemId, result[0].TransactionItems[0].Id)

	err = suite.sqlmock.ExpectationsWereMet()
	suite.NoError(err)
}

func (suite *TransactionsRepositoryTestSuite) TestGetAllWithSortByDesc() {
	limit := 3
	page := 0
	sortBy := "created_at"
	desc := true

	id := uuid.New()
	userId := uuid.New()
	grandTotal := 10.59
	status := string(constants.TransactionStatusConfirmed)
	invoiceNumber := "abcd"

	itemId := uuid.New()
	bookId := uuid.New()
	quantity := int64(1)
	price := 10.59

	// Expect count transactions
	countQuery := regexp.QuoteMeta(`SELECT count(*) FROM "transactions" WHERE user_id = $1`)
	suite.sqlmock.ExpectQuery(countQuery).
		WithArgs(userId).
		WillReturnRows(suite.sqlmock.NewRows([]string{"count"}).AddRow(1))

	// Expect select transactions table
	selectQuery := regexp.QuoteMeta(`SELECT * FROM "transactions" WHERE user_id = $1 AND "transactions"."deleted_at" IS NULL ORDER BY "created_at" DESC LIMIT $2`)
	suite.sqlmock.ExpectQuery(selectQuery).
		WithArgs(userId, limit).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_id", "grand_total", "status", "invoice_number"}).
				AddRow(id, userId, grandTotal, status, invoiceNumber),
		)

	// Expect select transaction_items table with Preload() functions.
	selectQuery = regexp.QuoteMeta(`SELECT * FROM "transaction_items" WHERE "transaction_items"."transaction_id" = $1`)
	suite.sqlmock.ExpectQuery(selectQuery).
		WithArgs(id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "transaction_id", "book_id", "quantity", "price"}).
				AddRow(itemId, id, bookId, quantity, price),
		)

	result, size, err := suite.repository.GetAll(userId.String(), limit, page, sortBy, desc)
	suite.Nil(err)
	suite.NotNil(result)
	suite.Equal(int64(1), size)
	suite.Equal(userId, result[0].UserId)
	suite.Equal(id, result[0].Id)
	suite.Equal(grandTotal, result[0].GrandTotal)
	suite.Equal(status, result[0].Status)
	suite.Equal(invoiceNumber, result[0].InvoiceNumber)
	suite.Equal(quantity, result[0].TransactionItems[0].Quantity)
	suite.Equal(price, result[0].TransactionItems[0].Price)
	suite.Equal(bookId, result[0].TransactionItems[0].BookId)
	suite.Equal(itemId, result[0].TransactionItems[0].Id)

	err = suite.sqlmock.ExpectationsWereMet()
	suite.NoError(err)
}

func (suite *TransactionsRepositoryTestSuite) TestGetAllWithSortBy_InvalidColumn() {
	limit := 3
	page := 0
	sortBy := "invalid"
	desc := false
	userId := uuid.New()

	// Expect count transactions
	countQuery := regexp.QuoteMeta(`SELECT count(*) FROM "transactions" WHERE user_id = $1`)
	suite.sqlmock.ExpectQuery(countQuery).
		WithArgs(userId).
		WillReturnRows(suite.sqlmock.NewRows([]string{"count"}).AddRow(1))

	result, size, err := suite.repository.GetAll(userId.String(), limit, page, sortBy, desc)
	suite.NotNil(err)
	suite.Empty(result)
	suite.Equal(int64(0), size)

	err = suite.sqlmock.ExpectationsWereMet()
	suite.NoError(err)
}

func (suite *TransactionsRepositoryTestSuite) TestGetByInvoiceNumber() {
	id := uuid.New()
	userId := uuid.New()
	grandTotal := 10.59
	status := string(constants.TransactionStatusConfirmed)
	invoiceNumber := "abcd"

	itemId := uuid.New()
	bookId := uuid.New()
	quantity := int64(1)
	price := 10.59

	// Expect select transactions table
	selectQuery := regexp.QuoteMeta(`SELECT * FROM "transactions" WHERE (user_id = $1 AND invoice_number = $2)`)
	suite.sqlmock.ExpectQuery(selectQuery).
		WithArgs(userId, invoiceNumber, 1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_id", "grand_total", "status", "invoice_number"}).
				AddRow(id, userId, grandTotal, status, invoiceNumber),
		)

	// Expect select transaction_items table with Preload() functions.
	selectQuery = regexp.QuoteMeta(`SELECT * FROM "transaction_items" WHERE "transaction_items"."transaction_id" = $1`)
	suite.sqlmock.ExpectQuery(selectQuery).
		WithArgs(id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "transaction_id", "book_id", "quantity", "price"}).
				AddRow(itemId, id, bookId, quantity, price),
		)

	result, err := suite.repository.GetByInvoiceNumber(userId.String(), invoiceNumber)
	suite.Nil(err)
	suite.NotNil(result)
	suite.Equal(userId, result.UserId)
	suite.Equal(id, result.Id)
	suite.Equal(grandTotal, result.GrandTotal)
	suite.Equal(status, result.Status)
	suite.Equal(invoiceNumber, result.InvoiceNumber)
	suite.Equal(quantity, result.TransactionItems[0].Quantity)
	suite.Equal(price, result.TransactionItems[0].Price)
	suite.Equal(bookId, result.TransactionItems[0].BookId)
	suite.Equal(itemId, result.TransactionItems[0].Id)

	err = suite.sqlmock.ExpectationsWereMet()
	suite.NoError(err)
}

func (suite *TransactionsRepositoryTestSuite) TestGetByInvoiceNumber_Error() {
	userId := uuid.New()
	invoiceNumber := "abcd"

	err := base_error.New("error")

	// Expect select transactions table
	selectQuery := regexp.QuoteMeta(`SELECT * FROM "transactions" WHERE (user_id = $1 AND invoice_number = $2)`)
	suite.sqlmock.ExpectQuery(selectQuery).
		WithArgs(userId, invoiceNumber, 1).
		WillReturnError(err)

	result, err := suite.repository.GetByInvoiceNumber(userId.String(), invoiceNumber)
	suite.NotNil(err)
	suite.Empty(result)

	err = suite.sqlmock.ExpectationsWereMet()
	suite.NoError(err)
}

func (suite *TransactionsRepositoryTestSuite) TestGetId() {
	id := uuid.New()
	userId := uuid.New()
	grandTotal := 10.59
	status := string(constants.TransactionStatusConfirmed)
	invoiceNumber := "abcd"

	itemId := uuid.New()
	bookId := uuid.New()
	quantity := int64(1)
	price := 10.59

	// Expect select transactions table
	selectQuery := regexp.QuoteMeta(`SELECT * FROM "transactions" WHERE (user_id = $1 AND id = $2)`)
	suite.sqlmock.ExpectQuery(selectQuery).
		WithArgs(userId, id, 1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user_id", "grand_total", "status", "invoice_number"}).
				AddRow(id, userId, grandTotal, status, invoiceNumber),
		)

	// Expect select transaction_items table with Preload() functions.
	selectQuery = regexp.QuoteMeta(`SELECT * FROM "transaction_items" WHERE "transaction_items"."transaction_id" = $1`)
	suite.sqlmock.ExpectQuery(selectQuery).
		WithArgs(id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "transaction_id", "book_id", "quantity", "price"}).
				AddRow(itemId, id, bookId, quantity, price),
		)

	result, err := suite.repository.GetById(userId.String(), id.String())
	suite.Nil(err)
	suite.NotNil(result)
	suite.Equal(userId, result.UserId)
	suite.Equal(id, result.Id)
	suite.Equal(grandTotal, result.GrandTotal)
	suite.Equal(status, result.Status)
	suite.Equal(invoiceNumber, result.InvoiceNumber)
	suite.Equal(quantity, result.TransactionItems[0].Quantity)
	suite.Equal(price, result.TransactionItems[0].Price)
	suite.Equal(bookId, result.TransactionItems[0].BookId)
	suite.Equal(itemId, result.TransactionItems[0].Id)

	err = suite.sqlmock.ExpectationsWereMet()
	suite.NoError(err)
}

func (suite *TransactionsRepositoryTestSuite) TestGetId_Error() {
	id := uuid.New()
	userId := uuid.New()

	err := base_error.New("error")

	// Expect select transactions table
	selectQuery := regexp.QuoteMeta(`SELECT * FROM "transactions" WHERE (user_id = $1 AND id = $2)`)
	suite.sqlmock.ExpectQuery(selectQuery).
		WithArgs(userId, id, 1).
		WillReturnError(err)

	result, err := suite.repository.GetById(userId.String(), id.String())
	suite.NotNil(err)
	suite.Empty(result)

	err = suite.sqlmock.ExpectationsWereMet()
	suite.NoError(err)
}

func (suite *TransactionsRepositoryTestSuite) TestCreate() {
	transactionId := uuid.New()

	input := models.Transactions{
		Id:            transactionId,
		UserId:        uuid.New(),
		GrandTotal:    10.59,
		InvoiceNumber: "abcd",
		Status:        string(constants.TransactionStatusConfirmed),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		TransactionItems: []models.TransactionItems{
			{
				Id:            uuid.New(),
				TransactionId: transactionId,
				BookId:        uuid.New(),
				Quantity:      1,
				Price:         10.59,
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
			},
		},
	}
	suite.sqlmock.ExpectBegin()
	suite.sqlmock.ExpectExec(`INSERT INTO "transactions"`).
		WithArgs(
			input.Id,
			input.UserId,
			input.GrandTotal,
			input.Status,
			input.InvoiceNumber,
			input.CreatedAt,
			input.UpdatedAt,
			input.DeletedAt,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.sqlmock.ExpectExec(`INSERT INTO "transaction_items"`).
		WithArgs(
			input.TransactionItems[0].Id,
			input.TransactionItems[0].TransactionId,
			input.TransactionItems[0].BookId,
			input.TransactionItems[0].Quantity,
			input.TransactionItems[0].Price,
			input.TransactionItems[0].CreatedAt,
			input.TransactionItems[0].UpdatedAt,
			input.TransactionItems[0].DeletedAt,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.sqlmock.ExpectCommit()
	result, err := suite.repository.Create(input)
	suite.Nil(err)
	suite.NotNil(result)

	err = suite.sqlmock.ExpectationsWereMet()
	suite.NoError(err)
}

func (suite *TransactionsRepositoryTestSuite) TestCreate_Error() {
	transactionId := uuid.New()

	input := models.Transactions{
		Id:            transactionId,
		UserId:        uuid.New(),
		GrandTotal:    10.59,
		InvoiceNumber: "abcd",
		Status:        string(constants.TransactionStatusConfirmed),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	suite.sqlmock.ExpectBegin()
	suite.sqlmock.ExpectExec(`INSERT INTO "transactions"`).
		WithArgs(
			input.Id,
			input.UserId,
			input.GrandTotal,
			input.Status,
			input.InvoiceNumber,
			input.CreatedAt,
			input.UpdatedAt,
			input.DeletedAt,
		).
		WillReturnResult(sqlmock.NewResult(0, 0))
	suite.sqlmock.ExpectCommit()
	result, err := suite.repository.Create(input)
	suite.NotNil(err)
	suite.Empty(result)

	err = suite.sqlmock.ExpectationsWereMet()
	suite.NoError(err)
}

func (suite *TransactionsRepositoryTestSuite) TestGenerateInvoiceCounter() {
	suite.sqlmock.ExpectExec("SELECT pg_advisory_lock").
		WithArgs(constants.LockKeyInvoiceNumber).
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.sqlmock.ExpectQuery("SELECT nextval").
		WithArgs(`invoice_counter_seq`).
		WillReturnRows(sqlmock.NewRows([]string{"nextval"}).AddRow(123))
	suite.sqlmock.ExpectExec("SELECT pg_advisory_unlock").
		WithArgs(constants.LockKeyInvoiceNumber).
		WillReturnResult(sqlmock.NewResult(1, 1))

	invoiceCounter, err := suite.repository.GenerateInvoiceCounter()
	suite.Nil(err)
	suite.NotEmpty(invoiceCounter)

	err = suite.sqlmock.ExpectationsWereMet()
	suite.NoError(err)
}
