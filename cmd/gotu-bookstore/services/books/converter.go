package books

import (
	contracts "gotu-bookstore/cmd/gotu-bookstore/contracts/books"
	"gotu-bookstore/cmd/gotu-bookstore/contracts/dto"
	"gotu-bookstore/cmd/gotu-bookstore/models"
	"gotu-bookstore/pkg/utils"
	"time"
)

func ConvertModelsIntoDTO(inputArray []models.Books) (contracts.GetBooksResponse, error) {
	var outputArray = make([]dto.BooksDTO, 0)
	for _, input := range inputArray {
		bookDTO, err := ConvertModelsIntoBookDetailsDTO(input)
		if err != nil {
			return nil, err
		}

		outputArray = append(outputArray, *bookDTO)
	}
	return outputArray, nil
}

func ConvertModelsIntoBookDetailsDTO(input models.Books) (*dto.BooksDTO, error) {
	bookDTO := dto.BooksDTO{
		Id:          input.Id.String(),
		Description: input.Description,
		PublishDate: input.PublishDate.Format(time.RFC3339Nano),
		ImageUrl:    input.ImageUrl,
		Page:        utils.Int64ToString(input.Page),
		Title:       input.Title,
		Price:       utils.FloatToString(input.Price),
		Language:    input.Language,
		Publisher:   input.Publisher,
		Author:      input.Author,
		Category:    input.Category,
		Isbn:        input.Isbn,
		Rating:      utils.FloatToString(input.Rating),
		CreatedAt:   input.CreatedAt.Format(time.RFC3339Nano),
		UpdatedAt:   input.UpdatedAt.Format(time.RFC3339Nano),
	}

	return &bookDTO, nil
}
