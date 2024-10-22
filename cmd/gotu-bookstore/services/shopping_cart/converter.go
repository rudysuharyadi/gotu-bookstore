package shopping_cart

import (
	"gotu-bookstore/cmd/gotu-bookstore/contracts/dto"
	contracts "gotu-bookstore/cmd/gotu-bookstore/contracts/shopping_cart"
	"gotu-bookstore/cmd/gotu-bookstore/models"
	"gotu-bookstore/pkg/utils"
)

func ConvertToShoppingCartResponse(cartItems []models.ShoppingCarts) (*contracts.ShoppingCartResponse, error) {
	var shoppingCartItemDTOs = make([]dto.ShoppingCartItemDTO, 0)
	grandTotal := 0.0
	for _, cartItem := range cartItems {
		shoppingCartItemDTOs = append(shoppingCartItemDTOs, dto.ShoppingCartItemDTO{
			BookId:   cartItem.BookId.String(),
			Quantity: cartItem.Quantity,
			Price:    utils.FloatToString(cartItem.Book.Price),
		})
		grandTotal += cartItem.Book.Price
	}

	return &contracts.ShoppingCartResponse{
		GrandTotal: utils.FloatToString(grandTotal),
		Items:      shoppingCartItemDTOs,
	}, nil
}
