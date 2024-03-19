package user_contracts_infra

import "github.com/junioralcant/api-stores-go/domain/models"

type IUserDeleteRepository interface {
	UserDeleteRepo(id string) (*models.User, error)
}
