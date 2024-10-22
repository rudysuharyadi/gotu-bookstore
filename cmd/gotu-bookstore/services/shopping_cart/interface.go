package shopping_cart

import "gotu-bookstore/cmd/gotu-bookstore/models"

type ShoppingCartRepoInterface interface {
	GetByUserId(userId string) ([]models.ShoppingCarts, error)
	GetItemByBookId(bookId string, userId string) (*models.ShoppingCarts, error)
	AddItemToCart(input models.ShoppingCarts) (*models.ShoppingCarts, error)
	UpdateByBookId(bookId string, quantity int64, userId string) error
	DeleteByBookId(bookId string, userId string) error
	ClearShoppingCart(userId string) error
}

type UserRepoInterface interface {
	GetById(id string) (*models.Users, error)
}

type BookRepoInterface interface {
	GetById(id string) (*models.Books, error)
}

type TransactionRepoInterface interface {
	Create(input models.Transactions) (*models.Transactions, error)
	GenerateInvoiceCounter() (int64, error)
}
