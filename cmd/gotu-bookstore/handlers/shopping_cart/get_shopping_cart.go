package shopping_cart

import (
	contracts "gotu-bookstore/cmd/gotu-bookstore/contracts/shopping_cart"
	"gotu-bookstore/pkg/utils"
)

type GetShoppingCartServiceInterface interface {
	ProcessingGetShoppingCart() (*contracts.ShoppingCartResponse, error)
}

type GetShoppingCartHandler struct {
	getShoppingCartService GetShoppingCartServiceInterface
	utils.CommonContext
}

func NewGetShoppingCartHandler(context utils.CommonContext, getShoppingCartService GetShoppingCartServiceInterface) GetShoppingCartHandler {
	return GetShoppingCartHandler{
		getShoppingCartService: getShoppingCartService,
		CommonContext:          context,
	}
}

func (h GetShoppingCartHandler) ProcessingGetShoppingCart() {
	result, err := h.getShoppingCartService.ProcessingGetShoppingCart()
	if err != nil {
		h.HandleFailure(err)
		return
	}

	h.HandleSuccess(result)
}
