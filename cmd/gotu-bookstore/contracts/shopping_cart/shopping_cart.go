package shopping_cart

import "gotu-bookstore/cmd/gotu-bookstore/contracts/dto"

type ShoppingCartResponse struct {
	GrandTotal string                    `json:"grand_total"`
	Items      []dto.ShoppingCartItemDTO `json:"items"`
}
