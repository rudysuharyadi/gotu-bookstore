package books

import (
	"gotu-bookstore/cmd/gotu-bookstore/constants"
	contracts "gotu-bookstore/cmd/gotu-bookstore/contracts/books"
	"gotu-bookstore/pkg/resfmt/base_error"
	"gotu-bookstore/pkg/utils"
)

type GetBooksServiceInterface interface {
	ProcessingGetBooks(request contracts.GetBooksRequest) (*contracts.GetBooksResponse, map[string]interface{}, error)
}

type GetBooksHandler struct {
	getBooksService GetBooksServiceInterface
	utils.CommonContext
}

func NewGetBooksHandler(context utils.CommonContext, getBooksService GetBooksServiceInterface) GetBooksHandler {
	return GetBooksHandler{
		getBooksService: getBooksService,
		CommonContext:   context,
	}
}

func (h GetBooksHandler) ProcessingGetBooks() {
	var request contracts.GetBooksRequest
	if err := h.GinContext.BindQuery(&request); err != nil {
		h.LogError(err)
		h.HandleFailure(base_error.NewBadRequestError(constants.IC0001))
		return
	}

	result, metadata, err := h.getBooksService.ProcessingGetBooks(request)
	if err != nil {
		h.HandleFailure(err)
		return
	}

	h.HandleSuccessWithMetadata(result, metadata)
}
