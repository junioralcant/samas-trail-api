package mocks

import "github.com/junioralcant/api-stores-go/domain/models"

type UserCreateUseCaseSpy struct {
	Params models.User
}

func NewUserCreateUseCaseSpy() *UserCreateUseCaseSpy {
	return &UserCreateUseCaseSpy{}
}

func (u *UserCreateUseCaseSpy) UserCreate(user models.User) *models.User {
	u.Params = user
	return &user
}
