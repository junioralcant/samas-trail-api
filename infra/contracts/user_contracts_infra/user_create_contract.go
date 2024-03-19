package user_contracts_infra

import "github.com/junioralcant/api-stores-go/domain/models"

type IUserCreateRepository interface {
	UserCreateRepo(user models.User) *models.User
}
