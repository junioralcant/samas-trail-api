package user_factories

import (
	"github.com/junioralcant/api-stores-go/data/usecases/user_usecase"
	"github.com/junioralcant/api-stores-go/infra/repositories/user_repository"
	"github.com/junioralcant/api-stores-go/presentation/controller/user_controller"
)

func UserUpdateControllerFactory() *user_controller.UserUpdateController {
	repo := user_repository.NewUserUpdateRepository()
	useCase := user_usecase.NewUserUpdateUseCase(repo)
	userUpdateController := user_controller.NewUserUpdateController(useCase)

	return userUpdateController
}
