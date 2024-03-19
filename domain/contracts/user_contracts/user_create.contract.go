package user_contracts

import "github.com/junioralcant/api-stores-go/domain/models"

type IUserCreate interface {
	UserCreate(user models.User) *models.User
}
