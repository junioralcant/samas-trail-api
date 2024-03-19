package user_contracts_infra

import "github.com/junioralcant/api-stores-go/domain/models"

type IUserRepository interface {
	ListAll() []models.User

	Create(user models.User) *models.User

	Update(id string, user models.User) (*models.User, error)

	Delete(id string) (*models.User, error)
}
