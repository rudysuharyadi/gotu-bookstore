package books

import (
	"gotu-bookstore/cmd/gotu-bookstore/constants"
	contracts "gotu-bookstore/cmd/gotu-bookstore/contracts/books"
	"gotu-bookstore/pkg/resfmt/base_error"
	"gotu-bookstore/pkg/utils"
)

type GetBookDetailsServiceInterface interface {
	ProcessingGetBookDetails(
		request contracts.GetBookDetailsRequest,
	) (*contracts.GetBookDetailsResponse, error)
}

type GetBookDetailsHandler struct {
	getBookDetailsService GetBookDetailsServiceInterface
	utils.CommonContext
}

func NewGetBookDetailsHandler(context utils.CommonContext, getBookDetailsService GetBookDetailsServiceInterface) GetBookDetailsHandler {
	return GetBookDetailsHandler{
		getBookDetailsService: getBookDetailsService,
		CommonContext:         context,
	}
}

func (h GetBookDetailsHandler) ProcessingGetBookDetails() {
	var request contracts.GetBookDetailsRequest
	if err := h.GinContext.BindUri(&request); err != nil {
		h.LogError(err)
		h.HandleFailure(base_error.NewBadRequestError(constants.IC0001))
		return
	}

	result, err := h.getBookDetailsService.ProcessingGetBookDetails(request)
	if err != nil {
		h.HandleFailure(err)
		return
	}

	h.HandleSuccess(result)
}
