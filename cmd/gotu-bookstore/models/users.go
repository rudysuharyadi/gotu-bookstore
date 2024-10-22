package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	Id        uuid.UUID `gorm:"primaryKey;unique"`
	Name      string
	Email     string
	Password  string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (n Users) TableName() string {
	return "users"
}
