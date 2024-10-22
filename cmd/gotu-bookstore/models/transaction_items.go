package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionItems struct {
	Id            uuid.UUID `gorm:"primaryKey;unique"`
	TransactionId uuid.UUID
	BookId        uuid.UUID
	Quantity      int64
	Price         float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

func (n TransactionItems) TableName() string {
	return "transaction_items"
}
