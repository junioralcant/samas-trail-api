package user_contracts

import "github.com/junioralcant/api-stores-go/domain/models"

type IUserDelete interface {
	UserDelete(id string) (*models.User, error)
}
