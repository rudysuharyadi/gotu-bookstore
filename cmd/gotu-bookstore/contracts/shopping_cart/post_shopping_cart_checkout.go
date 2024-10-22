package shopping_cart
import (
	"gotu-bookstore/cmd/gotu-bookstore/contracts/dto"
)

/*

*/
type PostShoppingCartCheckoutRequest struct {
}

/*
{
    "data": {
        "invoice_number": "inv-0001"
    },
    "status": "success"
}
*/
type PostShoppingCartCheckoutResponse struct {
	dto.ShoppingCartCheckoutDTO
}
