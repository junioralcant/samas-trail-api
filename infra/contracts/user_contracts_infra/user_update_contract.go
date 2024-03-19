package user_contracts_infra

import "github.com/junioralcant/api-stores-go/domain/models"

type IUserUpdateRepository interface {
	UserUpdateRepo(id string, user models.User) (*models.User, error)
}
