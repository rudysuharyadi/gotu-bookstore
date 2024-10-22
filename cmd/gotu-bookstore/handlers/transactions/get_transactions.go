package transactions

import (
	"gotu-bookstore/cmd/gotu-bookstore/constants"
	contracts "gotu-bookstore/cmd/gotu-bookstore/contracts/transactions"
	"gotu-bookstore/pkg/resfmt/base_error"
	"gotu-bookstore/pkg/utils"
)

type GetTransactionsServiceInterface interface {
	ProcessingGetTransactions(request contracts.GetTransactionsRequest) (*contracts.GetTransactionsResponse, map[string]interface{}, error)
}

type GetTransactionsHandler struct {
	getTransactionsService GetTransactionsServiceInterface
	utils.CommonContext
}

func NewGetTransactionsHandler(context utils.CommonContext, getTransactionsService GetTransactionsServiceInterface) GetTransactionsHandler {
	return GetTransactionsHandler{
		getTransactionsService: getTransactionsService,
		CommonContext:          context,
	}
}

func (h GetTransactionsHandler) ProcessingGetTransactions() {
	var request contracts.GetTransactionsRequest
	if err := h.GinContext.BindQuery(&request); err != nil {
		h.LogError(err)
		h.HandleFailure(base_error.NewBadRequestError(constants.IC0001))
		return
	}

	result, metadata, err := h.getTransactionsService.ProcessingGetTransactions(request)
	if err != nil {
		h.HandleFailure(err)
		return
	}

	h.HandleSuccessWithMetadata(result, metadata)
}
