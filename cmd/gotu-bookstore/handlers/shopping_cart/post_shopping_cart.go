package shopping_cart

import (
	"gotu-bookstore/cmd/gotu-bookstore/constants"
	contracts "gotu-bookstore/cmd/gotu-bookstore/contracts/shopping_cart"
	"gotu-bookstore/pkg/resfmt/base_error"
	"gotu-bookstore/pkg/utils"
)

type PostShoppingCartServiceInterface interface {
	ProcessingPostShoppingCart(request contracts.PostShoppingCartRequest) (*contracts.ShoppingCartResponse, error)
}

type PostShoppingCartHandler struct {
	postShoppingCartService PostShoppingCartServiceInterface
	utils.CommonContext
}

func NewPostShoppingCartHandler(context utils.CommonContext, postShoppingCartService PostShoppingCartServiceInterface) PostShoppingCartHandler {
	return PostShoppingCartHandler{
		postShoppingCartService: postShoppingCartService,
		CommonContext:           context,
	}
}

func (h PostShoppingCartHandler) ProcessingPostShoppingCart() {
	var request contracts.PostShoppingCartRequest
	if err := h.GinContext.BindJSON(&request); err != nil {
		h.LogError(err)
		h.HandleFailure(base_error.NewBadRequestError(constants.IC0001))
		return
	}

	result, err := h.postShoppingCartService.ProcessingPostShoppingCart(request)
	if err != nil {
		h.HandleFailure(err)
		return
	}

	h.HandleSuccess(result)
}
