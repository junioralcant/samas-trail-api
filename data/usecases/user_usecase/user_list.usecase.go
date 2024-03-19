package user_usecase

import (
	"github.com/junioralcant/api-stores-go/domain/models"
	"github.com/junioralcant/api-stores-go/infra/contracts/user_contracts_infra"
)

type UserListUseCase struct {
	Repo user_contracts_infra.IUserListAllRepository
}

func NewUserListUseCase(repo user_contracts_infra.IUserListAllRepository) *UserListUseCase {
	return &UserListUseCase{Repo: repo}
}

func (u *UserListUseCase) UserFindAll() []models.User {
	users := u.Repo.UserListAllRepo()

	return users
}
