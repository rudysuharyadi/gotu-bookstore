package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShoppingCarts struct {
	Id        uuid.UUID `gorm:"primaryKey;unique"`
	UserId    uuid.UUID
	BookId    uuid.UUID
	Quantity  int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Book      *Books
}

func (n ShoppingCarts) TableName() string {
	return "shopping_carts"
}
