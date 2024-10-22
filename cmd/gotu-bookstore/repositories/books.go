package repositories

import (
	"fmt"
	"gotu-bookstore/cmd/gotu-bookstore/models"
	"gotu-bookstore/pkg/database"
	"gotu-bookstore/pkg/resfmt/base_error"
	"slices"
	"strings"

	"gorm.io/gorm/clause"
)

type BooksRepository struct {
	dbInstance *database.Database
}

func NewBooksRepository(dbInstance *database.Database) BooksRepository {
	return BooksRepository{
		dbInstance: dbInstance,
	}
}

func (n BooksRepository) GetAll(limit int, page int, sortBy string, desc bool, keyword string) ([]models.Books, int64, error) {
	var results []models.Books
	var count int64
	db := n.dbInstance.DB

	// Base query
	query := db.Model(&models.Books{})

	if len(keyword) > 0 {
		keyword = strings.TrimSpace(keyword)
		keyword = fmt.Sprintf("%%%s%%", strings.ToLower(keyword))

		query = query.Where(
			db.Where("LOWER(title) LIKE ?", keyword).
				Or("LOWER(author) LIKE ?", keyword).
				Or("LOWER(description) LIKE ?", keyword).
				Or("LOWER(isbn) LIKE ?", keyword),
		)
	}

	err := query.Count(&count).Error
	if err != nil {
		return results, 0, err
	}

	if len(sortBy) > 0 {
		validColumn := models.Books{}.ValidSortColumn()
		if !slices.Contains(validColumn, sortBy) {
			err = base_error.New("Invalid sortBy column")
			return results, 0, err
		}

		query = query.Order(clause.OrderByColumn{
			Column: clause.Column{Name: sortBy},
			Desc:   desc,
		})
	}

	err = query.Limit(limit).Offset(page * limit).Find(&results).Error
	return results, count, err
}

func (n BooksRepository) GetById(id string) (*models.Books, error) {
	var results *models.Books
	err := n.dbInstance.DB.Where("id = ?", id).First(&results).Error
	return results, err
}
