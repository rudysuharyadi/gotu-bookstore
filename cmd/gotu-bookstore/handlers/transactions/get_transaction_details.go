package transactions

import (
	"gotu-bookstore/cmd/gotu-bookstore/constants"
	contracts "gotu-bookstore/cmd/gotu-bookstore/contracts/transactions"
	"gotu-bookstore/pkg/resfmt/base_error"
	"gotu-bookstore/pkg/utils"
)

type GetTransactionDetailsServiceInterface interface {
	ProcessingGetTransactionDetails(
		request contracts.GetTransactionDetailsRequest,
	) (*contracts.GetTransactionDetailsResponse, error)
}

type GetTransactionDetailsHandler struct {
	getTransactionDetailsService GetTransactionDetailsServiceInterface
	utils.CommonContext
}

func NewGetTransactionDetailsHandler(
	context utils.CommonContext,
	getTransactionDetailsService GetTransactionDetailsServiceInterface,
) GetTransactionDetailsHandler {
	return GetTransactionDetailsHandler{
		getTransactionDetailsService: getTransactionDetailsService,
		CommonContext:                context,
	}
}

func (h GetTransactionDetailsHandler) ProcessingGetTransactionDetails() {
	var request contracts.GetTransactionDetailsRequest
	if err := h.GinContext.BindUri(&request); err != nil {
		h.LogError(err)
		h.HandleFailure(base_error.NewBadRequestError(constants.IC0001))
		return
	}

	result, err := h.getTransactionDetailsService.ProcessingGetTransactionDetails(request)
	if err != nil {
		h.HandleFailure(err)
		return
	}

	h.HandleSuccess(result)
}
