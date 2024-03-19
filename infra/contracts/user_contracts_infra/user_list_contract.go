package user_contracts_infra

import "github.com/junioralcant/api-stores-go/domain/models"

type IUserListAllRepository interface {
	UserListAllRepo() []models.User
}
