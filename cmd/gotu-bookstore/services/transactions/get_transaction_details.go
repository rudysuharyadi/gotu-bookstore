package transactions

import (
	"gotu-bookstore/cmd/gotu-bookstore/constants"
	contracts "gotu-bookstore/cmd/gotu-bookstore/contracts/transactions"
	"gotu-bookstore/pkg/resfmt/base_error"
	"gotu-bookstore/pkg/utils"
)

type GetTransactionDetailsService struct {
	transactionsRepo TransactionRepoInterface
	userRepo         UserRepoInterface
	utils.CommonContext
}

func NewGetTransactionDetailsService(
	context utils.CommonContext,
	transactionsRepo TransactionRepoInterface,
	userRepo UserRepoInterface,
) GetTransactionDetailsService {
	return GetTransactionDetailsService{
		transactionsRepo: transactionsRepo,
		userRepo:         userRepo,
		CommonContext:    context,
	}
}

func (s GetTransactionDetailsService) ProcessingGetTransactionDetails(
	request contracts.GetTransactionDetailsRequest,
) (*contracts.GetTransactionDetailsResponse, error) {
	// Get session data
	session, err := s.GetSession()
	if err != nil {
		s.LogError(err)
		return nil, base_error.NewUnauthorizedError(constants.IC0006)
	}

	// Get user By ID from session
	user, err := s.userRepo.GetById(session.Id)
	if err != nil {
		s.LogDebug(err)
		return nil, base_error.NewInternalError(constants.IC0011)
	}
	if user == nil {
		return nil, base_error.NewUnauthorizedError(constants.IC0006)
	}

	transaction, err := s.transactionsRepo.GetById(user.Id.String(), request.TransactionId)
	if err != nil {
		s.LogError(err)
		return nil, base_error.NewInternalError(constants.IC0023)
	}

	transactionDTO, err := ConvertToTransactionDTO(*transaction)
	if err != nil {
		s.LogError(err)
		return nil, base_error.NewInternalError(constants.IC0023)
	}

	return &contracts.GetTransactionDetailsResponse{
		TransactionDTO: *transactionDTO,
	}, nil
}
