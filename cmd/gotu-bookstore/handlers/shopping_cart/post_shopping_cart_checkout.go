package shopping_cart

import (
	contracts "gotu-bookstore/cmd/gotu-bookstore/contracts/shopping_cart"
	"gotu-bookstore/pkg/utils"
)

type PostShoppingCartCheckoutServiceInterface interface {
	ProcessingPostShoppingCartCheckout() (*contracts.PostShoppingCartCheckoutResponse, error)
}

type PostShoppingCartCheckoutHandler struct {
	postShoppingCartCheckoutService PostShoppingCartCheckoutServiceInterface
	utils.CommonContext
}

func NewPostShoppingCartCheckoutHandler(context utils.CommonContext, postShoppingCartCheckoutService PostShoppingCartCheckoutServiceInterface) PostShoppingCartCheckoutHandler {
	return PostShoppingCartCheckoutHandler{
		postShoppingCartCheckoutService: postShoppingCartCheckoutService,
		CommonContext:                   context,
	}
}

func (h PostShoppingCartCheckoutHandler) ProcessingPostShoppingCartCheckout() {
	result, err := h.postShoppingCartCheckoutService.ProcessingPostShoppingCartCheckout()
	if err != nil {
		h.HandleFailure(err)
		return
	}

	h.HandleSuccess(result)
}
