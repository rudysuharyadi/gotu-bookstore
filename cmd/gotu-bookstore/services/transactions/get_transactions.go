package transactions

import (
	"gotu-bookstore/cmd/gotu-bookstore/constants"
	"gotu-bookstore/cmd/gotu-bookstore/contracts/pagination"
	contracts "gotu-bookstore/cmd/gotu-bookstore/contracts/transactions"
	"gotu-bookstore/cmd/gotu-bookstore/models"
	"gotu-bookstore/pkg/resfmt/base_error"
	"gotu-bookstore/pkg/utils"

	"github.com/mitchellh/mapstructure"
)

type GetTransactionsService struct {
	transactionsRepo TransactionRepoInterface
	userRepo         UserRepoInterface
	utils.CommonContext
}

func NewGetTransactionsService(
	context utils.CommonContext,
	transactionsRepo TransactionRepoInterface,
	userRepo UserRepoInterface,
) GetTransactionsService {
	return GetTransactionsService{
		transactionsRepo: transactionsRepo,
		userRepo:         userRepo,
		CommonContext:    context,
	}
}

func (s GetTransactionsService) ProcessingGetTransactions(request contracts.GetTransactionsRequest) (*contracts.GetTransactionsResponse, map[string]interface{}, error) {
	// Get session data
	session, err := s.GetSession()
	if err != nil {
		s.LogError(err)
		return nil, nil, base_error.NewUnauthorizedError(constants.IC0006)
	}

	// Get user By ID from session
	user, err := s.userRepo.GetById(session.Id)
	if err != nil {
		s.LogDebug(err)
		return nil, nil, base_error.NewInternalError(constants.IC0011)
	}
	if user == nil {
		return nil, nil, base_error.NewUnauthorizedError(constants.IC0006)
	}

	var results []models.Transactions
	var size int64
	var metadata = pagination.Metadata{
		Pagination: request.Pagination,
		Size:       0,
	}

	results, size, err = s.transactionsRepo.GetAll(user.Id.String(), request.Limit, request.Page, request.SortBy, request.Desc)
	metadata.Size = size
	if err != nil {
		s.LogDebug(err)
		return nil, nil, base_error.NewInternalError(constants.IC0023)
	}

	responses, err := ConvertToTransactionDTOs(results)
	if err != nil {
		s.LogError(err)
		return nil, nil, base_error.NewInternalError(constants.IC0023)
	}

	var metadataMap map[string]interface{}
	err = mapstructure.Decode(metadata, &metadataMap)
	if err != nil {
		s.LogError(err)
		return nil, nil, base_error.NewInternalError(constants.IC0023)
	}

	return (*contracts.GetTransactionsResponse)(&responses), metadataMap, nil
}
