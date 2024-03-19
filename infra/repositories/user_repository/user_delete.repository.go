package user_repository

import (
	"fmt"

	"github.com/junioralcant/api-stores-go/domain/models"
	"github.com/junioralcant/api-stores-go/infra/config"
)

type UserDeleteRepository struct {
}

func NewUserDeleteRepository() *UserDeleteRepository {
	return &UserDeleteRepository{}
}

func (r *UserDeleteRepository) UserDeleteRepo(id string) (*models.User, error) {
	userDelete := models.User{}

	if err := config.DB.First(&userDelete, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("error in search user: %+v", err)
	}

	if err := config.DB.Delete(&userDelete).Error; err != nil {
		return nil, fmt.Errorf("error in delete user: %+v", err)
	}

	return &userDelete, nil
}
