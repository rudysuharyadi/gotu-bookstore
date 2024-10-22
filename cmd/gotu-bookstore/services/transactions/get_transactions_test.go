package transactions_test

import (
	"net/http/httptest"
	"testing"
	"time"

	"gotu-bookstore/cmd/gotu-bookstore/constants"
	"gotu-bookstore/cmd/gotu-bookstore/contracts/pagination"
	contracts "gotu-bookstore/cmd/gotu-bookstore/contracts/transactions"
	"gotu-bookstore/cmd/gotu-bookstore/models"
	"gotu-bookstore/cmd/gotu-bookstore/services/transactions"
	mockery "gotu-bookstore/mocks/cmd/gotu-bookstore/services/transactions"
	"gotu-bookstore/pkg/auth/dto"
	"gotu-bookstore/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

func TestGetTransactionsServiceSuite(t *testing.T) {
	suite.Run(t, new(GetTransactionsServiceTestSuite))
}

type GetTransactionsServiceTestSuite struct {
	suite.Suite
	fakeTransactionRepo *mockery.TransactionRepoInterface
	fakeUserRepo        *mockery.UserRepoInterface
	service             transactions.GetTransactionsService
	context             utils.CommonContext
	sessionDTO          dto.SessionDTO
	user                models.Users
}

func (suite *GetTransactionsServiceTestSuite) SetupTest() {
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
	suite.fakeTransactionRepo = mockery.NewTransactionRepoInterface(suite.T())
	suite.fakeUserRepo = mockery.NewUserRepoInterface(suite.T())
	suite.service = transactions.NewGetTransactionsService(suite.context, suite.fakeTransactionRepo, suite.fakeUserRepo)
}

func (suite *GetTransactionsServiceTestSuite) TestGetAllTransactions() {
	request := contracts.GetTransactionsRequest{
		Pagination: pagination.Pagination{
			Limit:   10,
			Page:    0,
			SortBy:  "",
			Desc:    false,
			Keyword: "",
		},
	}

	transactions := GenerateFakerTransactions(3)

	suite.fakeUserRepo.EXPECT().GetById(suite.sessionDTO.Id).Return(&suite.user, nil)
	suite.fakeTransactionRepo.EXPECT().GetAll(suite.sessionDTO.Id, request.Limit, request.Page, request.SortBy, request.Desc).
		Return(transactions, int64(len(transactions)), nil)

	response, metadata, err := suite.service.ProcessingGetTransactions(request)
	suite.Nil(err)
	suite.NotNil(response)
	suite.NotNil(metadata)
}
