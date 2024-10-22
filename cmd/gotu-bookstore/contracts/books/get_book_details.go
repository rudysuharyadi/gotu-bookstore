package books

import "gotu-bookstore/cmd/gotu-bookstore/contracts/dto"

type GetBookDetailsRequest struct {
	BookId string `uri:"book_id"`
}

type GetBookDetailsResponse struct {
	dto.BooksDTO
}
