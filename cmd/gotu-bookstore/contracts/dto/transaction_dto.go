package dto

/*
	{
	    "created_at": "2023-02-01T01:01:01+0700",
	    "grand_total": "24.59",
	    "invoice_number": "inv-001",
	    "items": [
	        {
	            "book_id": "uuid",
	            "price": "10.00",
	            "quantity": 1
	        }
	    ],
	    "status": "CONFIRMED"
	}
*/
type TransactionDTO struct {
	Id            string               `json:"id"`
	Items         []TransactionItemDTO `json:"items"`
	InvoiceNumber string               `json:"invoice_number"`
	GrandTotal    string               `json:"grand_total"`
	Status        string               `json:"status"`
	CreatedAt     string               `json:"created_at"`
}
