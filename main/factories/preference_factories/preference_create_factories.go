package preference_factories

import (
	"github.com/junioralcant/api-stores-go/data/usecases/preference_usecase"
	preference_repository "github.com/junioralcant/api-stores-go/infra/repositories/preference"
	"github.com/junioralcant/api-stores-go/presentation/controller/preference_controller"
)

func PreferenceCreateControllerFactory() *preference_controller.PreferenceCreateController {
	repo := preference_repository.NewPreferenceRepository()
	useCaseCreate := preference_usecase.NewPreferenceCreateUseCase(repo)
	preferenceCreateController := preference_controller.NewPreferenceCreateController(useCaseCreate)

	return preferenceCreateController
}
