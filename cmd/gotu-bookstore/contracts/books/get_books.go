package books

import (
	"gotu-bookstore/cmd/gotu-bookstore/contracts/dto"
	"gotu-bookstore/cmd/gotu-bookstore/contracts/pagination"
)

/*
 */
type GetBooksRequest struct {
	pagination.Pagination
}

/*
	{
	    "data": [{
	        "author": "Sarah Andersen",
	        "category": "webcomic",
	        "created_at": "2023-02-01T01:01:01+0700",
	        "deleted_at": "2023-02-01T01:01:01+0700",
	        "description": "Sarah's Scribbles is a webcomic by Sarah Andersen started in 2011.",
	        "discount_amount": "1.89",
	        "discount_rate": "10%",
	        "id": "uuid",
	        "image_url": "",
	        "isbn": "9788833140940",
	        "language": "english",
	        "page": "150",
	        "price": "18.99",
	        "publish_date": "2021",
	        "publisher": "Tapas Media",
	        "rating": "4.5",
	        "title": "Sarah's Scribbles",
	        "updated_at": "2023-02-01T01:01:01+0700"
	    }],
	    "status": "success"
	}
*/
type GetBooksResponse []dto.BooksDTO
