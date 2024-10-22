package transactions

import (
	"gotu-bookstore/cmd/gotu-bookstore/contracts/dto"
	"gotu-bookstore/cmd/gotu-bookstore/contracts/pagination"
)

/*
 */
type GetTransactionsRequest struct {
	pagination.Pagination
}

/*
	{
	    "data": [
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
	    ],
	    "status": "success"
	}
*/
type GetTransactionsResponse []dto.TransactionDTO
