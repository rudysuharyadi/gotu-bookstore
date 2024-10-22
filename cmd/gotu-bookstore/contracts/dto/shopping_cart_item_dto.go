package dto

/*
	{
	    "book_id": "uuid",
	    "quantity": 1
	}
*/
type ShoppingCartItemDTO struct {
	BookId   string `json:"book_id"`
	Quantity int64  `json:"quantity"`
	Price    string `json:"price"`
}
