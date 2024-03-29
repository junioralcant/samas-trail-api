package user_repository

import (
	"fmt"

	"github.com/junioralcant/api-stores-go/domain/models"
	"github.com/junioralcant/api-stores-go/infra/config"
)

type UserUpdateRepository struct {
}

func NewUserUpdateRepository() *UserUpdateRepository {
	return &UserUpdateRepository{}
}

func (r *UserUpdateRepository) UserUpdateRepo(id string, user models.User) (*models.User, error) {

	userUpdated := models.User{}

	if err := config.DB.First(&userUpdated, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("error in search user: %+v", err)
	}

	if user.ID != "" {
		userUpdated.ID = user.ID
	}

	if user.IdPreferences != "" {
		userUpdated.IdPreferences = user.IdPreferences
	}

	if user.Name != "" {
		userUpdated.Name = user.Name
	}

	if user.Email != "" {
		userUpdated.Email = user.Email
	}

	if user.Phone != "" {
		userUpdated.Phone = user.Phone
	}

	if user.CPF != "" {
		userUpdated.CPF = user.CPF
	}

	if user.City != "" {
		userUpdated.City = user.City
	}

	if user.Team != "" {
		userUpdated.Team = user.Team
	}

	if user.Street != "" {
		userUpdated.Street = user.Street
	}

	if user.PostalCode != "" {
		userUpdated.PostalCode = user.PostalCode
	}

	if user.Paid != nil {
		userUpdated.Paid = user.Paid
	}

	if err := config.DB.Save(&userUpdated).Error; err != nil {
		return nil, fmt.Errorf("error in update user: %+v", err)
	}

	return &userUpdated, nil
}
