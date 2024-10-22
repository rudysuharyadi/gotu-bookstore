package transactions

import "gotu-bookstore/cmd/gotu-bookstore/models"

type TransactionRepoInterface interface {
	GetAll(userId string, limit int, page int, sortBy string, desc bool) ([]models.Transactions, int64, error)
	GetByInvoiceNumber(userId string, invoiceNumber string) (*models.Transactions, error)
	GetById(userId string, id string) (*models.Transactions, error)
}

type UserRepoInterface interface {
	GetById(id string) (*models.Users, error)
}
