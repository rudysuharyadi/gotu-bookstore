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

func TestUsersRepositorySuite(t *testing.T) {
	suite.Run(t, new(UsersRepositoryTestSuite))
}

type UsersRepositoryTestSuite struct {
	sqlmock    sqlmock.Sqlmock
	repository repositories.UsersRepository
	suite.Suite
}

func (suite *UsersRepositoryTestSuite) SetupTest() {
	db, mock, err := database.NewMockDB()
	suite.Nil(err)
	suite.sqlmock = mock
	suite.repository = repositories.NewUsersRepository(db)
}

func (suite *UsersRepositoryTestSuite) TestGetById() {
	id := uuid.New()
	name := "Rudy Suharyadi"
	email := "rudy.suharyadi@gmail.com"
	password := "password"
	status := string(constants.UserStatusActive)

	selectQuery := regexp.QuoteMeta(`SELECT * FROM "users" WHERE id = $1`)
	suite.sqlmock.ExpectQuery(selectQuery).
		WithArgs(id, 1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "email", "password", "status"}).
				AddRow(id, name, email, password, status),
		)

	result, err := suite.repository.GetById(id.String())
	suite.Nil(err)
	suite.NotNil(result)
	suite.Equal(name, result.Name)
	suite.Equal(id, result.Id)
	suite.Equal(email, result.Email)
	suite.Equal(password, result.Password)
	suite.Equal(status, result.Status)

	err = suite.sqlmock.ExpectationsWereMet()
	suite.NoError(err)
}

func (suite *UsersRepositoryTestSuite) TestGetById_Error() {
	id := uuid.New()
	err := base_error.New("test")

	selectQuery := regexp.QuoteMeta(`SELECT * FROM "users" WHERE id = $1`)
	suite.sqlmock.ExpectQuery(selectQuery).
		WithArgs(id, 1).
		WillReturnError(err)

	result, err := suite.repository.GetById(id.String())
	suite.NotNil(err)
	suite.Empty(result)

	err = suite.sqlmock.ExpectationsWereMet()
	suite.NoError(err)
}

func (suite *UsersRepositoryTestSuite) TestGetByEmail() {
	id := uuid.New()
	name := "Rudy Suharyadi"
	email := "rudy.suharyadi@gmail.com"
	password := "password"
	status := string(constants.UserStatusActive)

	selectQuery := regexp.QuoteMeta(`SELECT * FROM "users" WHERE email = $1`)
	suite.sqlmock.ExpectQuery(selectQuery).
		WithArgs(email, 1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "email", "password", "status"}).
				AddRow(id, name, email, password, status),
		)

	result, err := suite.repository.GetByEmail(email)
	suite.Nil(err)
	suite.NotNil(result)
	suite.Equal(name, result.Name)
	suite.Equal(id, result.Id)
	suite.Equal(email, result.Email)
	suite.Equal(password, result.Password)
	suite.Equal(status, result.Status)

	err = suite.sqlmock.ExpectationsWereMet()
	suite.NoError(err)
}

func (suite *UsersRepositoryTestSuite) TestGetByEmail_Error() {
	email := "rudy.suharyadi@gmail.com"
	err := base_error.New("test")

	selectQuery := regexp.QuoteMeta(`SELECT * FROM "users" WHERE email = $1`)
	suite.sqlmock.ExpectQuery(selectQuery).
		WithArgs(email, 1).
		WillReturnError(err)

	result, err := suite.repository.GetByEmail(email)
	suite.NotNil(err)
	suite.Empty(result)

	err = suite.sqlmock.ExpectationsWereMet()
	suite.NoError(err)
}

func (suite *UsersRepositoryTestSuite) TestCreate() {
	input := models.Users{
		Id:        uuid.New(),
		Name:      "Rudy Suharyadi",
		Email:     "rudy.suharyadi@gmail.com",
		Password:  "pasword",
		Status:    string(constants.UserStatusActive),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	suite.sqlmock.ExpectBegin()
	suite.sqlmock.ExpectExec(`INSERT INTO "users"`).
		WithArgs(
			input.Id,
			input.Name,
			input.Email,
			input.Password,
			input.Status,
			input.CreatedAt,
			input.UpdatedAt,
			input.DeletedAt,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.sqlmock.ExpectCommit()
	result, err := suite.repository.Create(input)
	suite.Nil(err)
	suite.NotNil(result)

	err = suite.sqlmock.ExpectationsWereMet()
	suite.NoError(err)
}

func (suite *UsersRepositoryTestSuite) TestCreate_Error() {
	input := models.Users{
		Id:        uuid.New(),
		Name:      "Rudy Suharyadi",
		Email:     "rudy.suharyadi@gmail.com",
		Password:  "pasword",
		Status:    string(constants.UserStatusActive),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	suite.sqlmock.ExpectBegin()
	suite.sqlmock.ExpectExec(`INSERT INTO "users"`).
		WithArgs(
			input.Id,
			input.Name,
			input.Email,
			input.Password,
			input.Status,
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
