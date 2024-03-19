package user_factories

import (
	"github.com/junioralcant/api-stores-go/data/usecases/user_usecase"
	"github.com/junioralcant/api-stores-go/infra/repositories/user_repository"
	"github.com/junioralcant/api-stores-go/presentation/controller/user_controller"
)

func UserDeleteControllerFactory() *user_controller.UserDeleteController {
	repo := user_repository.NewUserDeleteRepository()
	useCaseCreate := user_usecase.NewUserDeleteUseCase(repo)
	userDeleteController := user_controller.NewUserDeleteController(useCaseCreate)

	return userDeleteController
}
