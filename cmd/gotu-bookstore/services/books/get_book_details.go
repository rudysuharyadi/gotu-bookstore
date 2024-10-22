package books

import (
	"gotu-bookstore/cmd/gotu-bookstore/constants"
	contracts "gotu-bookstore/cmd/gotu-bookstore/contracts/books"
	"gotu-bookstore/pkg/resfmt/base_error"
	"gotu-bookstore/pkg/utils"
)

type GetBookDetailsService struct {
	bookRepo BooksRepoInterface
	utils.CommonContext
}

func NewGetBookDetailsService(context utils.CommonContext, bookRepo BooksRepoInterface) GetBookDetailsService {
	return GetBookDetailsService{
		bookRepo:      bookRepo,
		CommonContext: context,
	}
}

func (s GetBookDetailsService) ProcessingGetBookDetails(
	request contracts.GetBookDetailsRequest,
) (*contracts.GetBookDetailsResponse, error) {

	book, err := s.bookRepo.GetById(request.BookId)
	if err != nil {
		s.LogDebug(err)
		return nil, base_error.NewInternalError(constants.IC0014)
	}

	response, err := ConvertModelsIntoBookDetailsDTO(*book)
	if err != nil {
		s.LogError(err)
		return nil, base_error.NewInternalError(constants.IC0014)
	}

	return &contracts.GetBookDetailsResponse{
		BooksDTO: *response,
	}, nil
}
