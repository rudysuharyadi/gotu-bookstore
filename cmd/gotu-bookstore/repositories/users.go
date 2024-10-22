package repositories

import (
	"gotu-bookstore/cmd/gotu-bookstore/models"
	"gotu-bookstore/pkg/database"
	"gotu-bookstore/pkg/resfmt/base_error"
)

type UsersRepository struct {
	dbInstance *database.Database
}

func NewUsersRepository(dbInstance *database.Database) UsersRepository {
	return UsersRepository{
		dbInstance: dbInstance,
	}
}

func (n UsersRepository) GetById(id string) (*models.Users, error) {
	var results *models.Users
	err := n.dbInstance.DB.Where("id = ?", id).First(&results).Error
	return results, err
}

func (n UsersRepository) GetByEmail(email string) (*models.Users, error) {
	var results *models.Users
	err := n.dbInstance.DB.Where("email = ?", email).First(&results).Error
	return results, err
}

func (n UsersRepository) Create(input models.Users) (*models.Users, error) {
	result := n.dbInstance.DB.Table(input.TableName()).Create(input)
	if result.Error == nil && result.RowsAffected == 0 {
		return nil, base_error.New("UsersRepository - Create - No data found.")
	}
	return &input, result.Error
}
