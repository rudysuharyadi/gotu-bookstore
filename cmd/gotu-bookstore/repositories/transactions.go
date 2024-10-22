package repositories

import (
	"gotu-bookstore/cmd/gotu-bookstore/constants"
	"gotu-bookstore/cmd/gotu-bookstore/models"
	"gotu-bookstore/pkg/database"
	"gotu-bookstore/pkg/resfmt/base_error"
	"slices"

	"gorm.io/gorm/clause"
)

type TransactionsRepository struct {
	dbInstance *database.Database
}

func NewTransactionsRepository(dbInstance *database.Database) TransactionsRepository {
	return TransactionsRepository{
		dbInstance: dbInstance,
	}
}

func (n TransactionsRepository) GetAll(userId string, limit int, page int, sortBy string, desc bool) ([]models.Transactions, int64, error) {
	var results []models.Transactions
	var count int64
	db := n.dbInstance.DB

	// Base query
	query := db.Model(&models.Transactions{}).Where("user_id = ?", userId)

	err := query.Count(&count).Error
	if err != nil {
		return results, 0, err
	}

	if len(sortBy) > 0 {
		validColumn := models.Transactions{}.ValidSortColumn()
		if !slices.Contains(validColumn, sortBy) {
			err = base_error.New("Invalid sortBy column")
			return results, 0, err
		}

		query = query.Order(clause.OrderByColumn{
			Column: clause.Column{Name: sortBy},
			Desc:   desc,
		})
	}

	err = query.Preload("TransactionItems").Limit(limit).Offset(page * limit).Find(&results).Error
	return results, count, err
}

func (n TransactionsRepository) GetByInvoiceNumber(userId, invoiceNumber string) (*models.Transactions, error) {
	var results *models.Transactions
	err := n.dbInstance.DB.Preload("TransactionItems").
		Where("user_id = ? AND invoice_number = ?", userId, invoiceNumber).
		First(&results).
		Error
	return results, err
}

func (n TransactionsRepository) GetById(userId, id string) (*models.Transactions, error) {
	var results *models.Transactions
	err := n.dbInstance.DB.Preload("TransactionItems").
		Where("user_id = ? AND id = ?", userId, id).
		First(&results).
		Error
	return results, err
}

func (n TransactionsRepository) Create(input models.Transactions) (*models.Transactions, error) {
	result := n.dbInstance.DB.Table(input.TableName()).Create(input)
	if result.Error == nil && result.RowsAffected == 0 {
		return nil, base_error.New("TransactionsRepository - Create - No data found.")
	}
	return &input, result.Error
}

func (n TransactionsRepository) GenerateInvoiceCounter() (int64, error) {
	lockKey := int64(constants.LockKeyInvoiceNumber)
	err := n.dbInstance.AcquireAdvisoryLock(lockKey)
	if err != nil {
		return 0, err
	}
	defer n.dbInstance.ReleaseAdvisoryLock(lockKey)

	var nextID int64
	err = n.dbInstance.DB.Raw("SELECT nextval(?)", "invoice_counter_seq").Scan(&nextID).Error
	if err != nil {
		return 0, err
	}

	return nextID, nil
}
