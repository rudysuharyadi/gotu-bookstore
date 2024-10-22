package books

import (
	"gotu-bookstore/cmd/gotu-bookstore/constants"
	contracts "gotu-bookstore/cmd/gotu-bookstore/contracts/books"
	"gotu-bookstore/cmd/gotu-bookstore/contracts/pagination"
	"gotu-bookstore/cmd/gotu-bookstore/models"
	"gotu-bookstore/pkg/resfmt/base_error"
	"gotu-bookstore/pkg/utils"

	"github.com/mitchellh/mapstructure"
)

type GetBooksService struct {
	bookRepo BooksRepoInterface
	utils.CommonContext
}

func NewGetBooksService(context utils.CommonContext, bookRepo BooksRepoInterface) GetBooksService {
	return GetBooksService{
		bookRepo:      bookRepo,
		CommonContext: context,
	}
}

func (s GetBooksService) ProcessingGetBooks(request contracts.GetBooksRequest) (*contracts.GetBooksResponse, map[string]interface{}, error) {
	var results []models.Books
	var size int64
	var err error
	var metadata = pagination.Metadata{
		Pagination: request.Pagination,
		Size:       0,
	}

	results, size, err = s.bookRepo.GetAll(request.Limit, request.Page, request.SortBy, request.Desc, request.Keyword)
	metadata.Size = size
	if err != nil {
		s.LogDebug(err)
		return nil, nil, base_error.NewInternalError(constants.IC0014)
	}

	responses, err := ConvertModelsIntoDTO(results)
	if err != nil {
		s.LogError(err)
		return nil, nil, base_error.NewInternalError(constants.IC0014)
	}

	var metadataMap map[string]interface{}
	err = mapstructure.Decode(metadata, &metadataMap)
	if err != nil {
		s.LogError(err)
		return nil, nil, base_error.NewInternalError(constants.IC0014)
	}

	return &responses, metadataMap, nil
}
