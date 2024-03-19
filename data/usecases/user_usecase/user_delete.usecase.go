package user_usecase

import (
	"github.com/junioralcant/api-stores-go/domain/models"
	"github.com/junioralcant/api-stores-go/infra/repositories/user_repository"
)

type UserDeleteUseCase struct {
	Repo *user_repository.UserDeleteRepository
}

func NewUserDeleteUseCase(repo *user_repository.UserDeleteRepository) *UserDeleteUseCase {
	return &UserDeleteUseCase{Repo: repo}
}

func (uc *UserDeleteUseCase) UserDelete(id string) (*models.User, error) {

	userDelete, err := uc.Repo.UserDeleteRepo(id)
	if err != nil {
		return nil, err
	}

	return userDelete, nil
}
