package shopping_cart

/*
	{
	    "book_id": "uuid",
	    "quantity": 1
	}
*/
type PutShoppingCartRequest struct {
	BookId   string `json:"book_id"`
	Quantity int    `json:"quantity"`
}
