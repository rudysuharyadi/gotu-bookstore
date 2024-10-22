package dto

/*
	{
	    "book_id": "uuid",
	    "price": "10.00",
	    "quantity": 1
	}
*/
type TransactionItemDTO struct {
	BookId   string `json:"book_id"`
	Quantity int64  `json:"quantity"`
	Price    string `json:"price"`
}
