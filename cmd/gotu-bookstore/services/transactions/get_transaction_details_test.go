package transactions_test

import (
	"net/http/httptest"
	"testing"
	"time"

	"gotu-bookstore/cmd/gotu-bookstore/constants"
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

func TestGetTransactionDetailsServiceSuite(t *testing.T) {
	suite.Run(t, new(GetTransactionDetailsServiceTestSuite))
}

type GetTransactionDetailsServiceTestSuite struct {
	suite.Suite
	fakeTransactionRepo *mockery.TransactionRepoInterface
	fakeUserRepo        *mockery.UserRepoInterface
	service             transactions.GetTransactionDetailsService
	context             utils.CommonContext
	sessionDTO          dto.SessionDTO
	user                models.Users
}

func (suite *GetTransactionDetailsServiceTestSuite) SetupTest() {
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
	suite.service = transactions.NewGetTransactionDetailsService(suite.context, suite.fakeTransactionRepo, suite.fakeUserRepo)
}

func (suite *GetTransactionDetailsServiceTestSuite) TestGetTransactionDetails() {
	transaction := GenerateFakerTransactions(1)[0]
	request := contracts.GetTransactionDetailsRequest{
		TransactionId: transaction.Id.String(),
	}

	suite.fakeUserRepo.EXPECT().GetById(suite.sessionDTO.Id).Return(&suite.user, nil)
	suite.fakeTransactionRepo.EXPECT().GetById(suite.sessionDTO.Id, transaction.Id.String()).
		Return(&transaction, nil)

	response, err := suite.service.ProcessingGetTransactionDetails(request)
	suite.Nil(err)
	suite.NotNil(response)
	suite.Equal(transaction.Id.String(), response.Id)
	suite.Equal(transaction.InvoiceNumber, response.InvoiceNumber)
	suite.Equal(utils.FloatToString(transaction.GrandTotal), response.GrandTotal)
	suite.Equal(transaction.Status, response.Status)
	suite.Equal(len(transaction.TransactionItems), len(response.Items))
}
