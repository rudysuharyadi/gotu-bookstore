package books

import "gotu-bookstore/cmd/gotu-bookstore/models"

type BooksRepoInterface interface {
	GetAll(limit int, page int, sortBy string, desc bool, keyword string) ([]models.Books, int64, error)
	GetById(id string) (*models.Books, error)
}
