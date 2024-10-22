package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transactions struct {
	Id               uuid.UUID `gorm:"primaryKey;unique"`
	UserId           uuid.UUID
	GrandTotal       float64
	Status           string
	InvoiceNumber    string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
	TransactionItems []TransactionItems `gorm:"foreignKey:TransactionId;references:Id"`
}

func (n Transactions) TableName() string {
	return "transactions"
}

func (n Transactions) ValidSortColumn() []string {
	return []string{"user_id", "status", "created_at"}
}
