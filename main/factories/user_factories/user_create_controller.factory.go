package user_factories

import (
	"github.com/junioralcant/api-stores-go/data/usecases/user_usecase"
	"github.com/junioralcant/api-stores-go/infra/repositories/user_repository"
	"github.com/junioralcant/api-stores-go/presentation/controller/user_controller"
)

func UserCreateControllerFactory() *user_controller.UserCreateController {
	repo := user_repository.NewUserCreateRepository()
	useCaseCreate := user_usecase.NewUserCreateUseCase(repo)
	userCreateController := user_controller.NewUserCreateController(useCaseCreate)

	return userCreateController
}
