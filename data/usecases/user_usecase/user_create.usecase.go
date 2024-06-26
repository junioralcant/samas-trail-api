package user_usecase

import (
	"github.com/junioralcant/api-stores-go/domain/models"
	"github.com/junioralcant/api-stores-go/infra/contracts/user_contracts_infra"
)

type UserCreateUseCase struct {
	Repo user_contracts_infra.IUserCreateRepository
}

func NewUserCreateUseCase(repo user_contracts_infra.IUserCreateRepository) *UserCreateUseCase {
	return &UserCreateUseCase{Repo: repo}
}

func (u *UserCreateUseCase) UserCreate(user models.User) (*models.User, error) {
	userCreated, err := u.Repo.UserCreateRepo(user)

	if err != nil {
		return nil, err
	}

	return userCreated, nil
}
