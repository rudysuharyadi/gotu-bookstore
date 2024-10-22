package repositories

import (
	"gotu-bookstore/cmd/gotu-bookstore/models"
	"gotu-bookstore/pkg/database"
	"gotu-bookstore/pkg/resfmt/base_error"
	"time"

	"gorm.io/gorm"
)

type ShoppingCartsRepository struct {
	dbInstance *database.Database
}

func NewShoppingCartsRepository(dbInstance *database.Database) ShoppingCartsRepository {
	return ShoppingCartsRepository{
		dbInstance: dbInstance,
	}
}

func (n ShoppingCartsRepository) GetByUserId(userId string) ([]models.ShoppingCarts, error) {
	var results []models.ShoppingCarts
	err := n.dbInstance.DB.Model(&models.ShoppingCarts{}).
		Preload("Book").
		Where("user_id = ?", userId).
		Find(&results).
		Error
	return results, err
}

func (n ShoppingCartsRepository) GetItemByBookId(bookId string, userId string) (*models.ShoppingCarts, error) {
	var results *models.ShoppingCarts
	err := n.dbInstance.DB.Where("book_id = ? AND user_id = ?", bookId, userId).First(&results).Error
	if base_error.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return results, err
}

func (n ShoppingCartsRepository) AddItemToCart(input models.ShoppingCarts) (*models.ShoppingCarts, error) {
	result := n.dbInstance.DB.Table(input.TableName()).Create(input)
	if result.Error == nil && result.RowsAffected == 0 {
		return nil, base_error.New("ShoppingCartsRepository - Create - No data found.")
	}
	return &input, result.Error
}

func (n ShoppingCartsRepository) UpdateByBookId(bookId string, quantity int64, userId string) error {
	result := n.dbInstance.DB.Model(&models.ShoppingCarts{}).
		Where("book_id = ? AND user_id = ?", bookId, userId).
		Updates(models.ShoppingCarts{
			Quantity:  quantity,
			UpdatedAt: time.Now(),
		})
	if result.RowsAffected == 0 {
		return base_error.New("ShoppingCartsRepository - UpdateByBookId - No data found.")
	}
	return result.Error
}

func (n ShoppingCartsRepository) DeleteByBookId(bookId string, userId string) error {
	result := n.dbInstance.DB.Where("book_id = ? AND user_id = ?", bookId, userId).Delete(&models.ShoppingCarts{})
	if result.RowsAffected == 0 {
		return base_error.New("ShoppingCartsRepository - DeleteById - No data found.")
	}
	return result.Error
}

func (n ShoppingCartsRepository) ClearShoppingCart(userId string) error {
	result := n.dbInstance.DB.Where("user_id = ?", userId).Delete(&models.ShoppingCarts{})
	if result.RowsAffected == 0 {
		return base_error.New("ShoppingCartsRepository - DeleteById - No data found.")
	}
	return result.Error
}
