package user_repository

import (
	"github.com/junioralcant/api-stores-go/domain/models"
	"github.com/junioralcant/api-stores-go/infra/config"
)

type UserListAllRepository struct {
}

func NewUserListRepository() *UserListAllRepository {
	return &UserListAllRepository{}
}

func (r *UserListAllRepository) UserListAllRepo() []models.User {
	users := []models.User{}

	if err := config.DB.Find(&users).Error; err != nil {
		config.Log.Errorf("error in list openings: %v", err)
	}

	return users
}
