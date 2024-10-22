package shopping_cart

import (
	contracts "gotu-bookstore/cmd/gotu-bookstore/contracts/shopping_cart"
	"gotu-bookstore/pkg/utils"
)

type PostShoppingCartClearServiceInterface interface {
	ProcessingPostShoppingCartClear() (*contracts.PostShoppingCartClearResponse, error)
}

type PostShoppingCartClearHandler struct {
	postShoppingCartClearService PostShoppingCartClearServiceInterface
	utils.CommonContext
}

func NewPostShoppingCartClearHandler(context utils.CommonContext, postShoppingCartClearService PostShoppingCartClearServiceInterface) PostShoppingCartClearHandler {
	return PostShoppingCartClearHandler{
		postShoppingCartClearService: postShoppingCartClearService,
		CommonContext:                context,
	}
}

func (h PostShoppingCartClearHandler) ProcessingPostShoppingCartClear() {
	result, err := h.postShoppingCartClearService.ProcessingPostShoppingCartClear()
	if err != nil {
		h.HandleFailure(err)
		return
	}

	h.HandleSuccess(result)
}
