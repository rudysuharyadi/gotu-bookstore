package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Books struct {
	Id          uuid.UUID `gorm:"primaryKey;unique"`
	Author      string
	Title       string
	Description string
	Category    string
	Publisher   string
	Price       float64
	Isbn        string
	Language    string
	PublishDate time.Time
	ImageUrl    string
	Page        int64
	Rating      float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

func (n Books) TableName() string {
	return "books"
}

func (n Books) ValidSortColumn() []string {
	return []string{"author", "title", "category", "publisher", "rating", "language", "isbn", "publish_date", "created_at"}
}
