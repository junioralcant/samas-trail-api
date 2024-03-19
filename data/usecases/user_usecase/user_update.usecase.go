package user_usecase

import (
	"github.com/junioralcant/api-stores-go/domain/models"
	"github.com/junioralcant/api-stores-go/infra/contracts/user_contracts_infra"
)

type UserUpdateUseCase struct {
	Repo user_contracts_infra.IUserUpdateRepository
}

func NewUserUpdateUseCase(repo user_contracts_infra.IUserUpdateRepository) *UserUpdateUseCase {
	return &UserUpdateUseCase{Repo: repo}
}

func (u *UserUpdateUseCase) UserUpdate(id string, user models.User) (*models.User, error) {
	userCreated, err := u.Repo.UserUpdateRepo(id, user)
	if err != nil {
		return nil, err
	}
	return userCreated, nil
}
