package preference_usecase

import (
	"github.com/junioralcant/api-stores-go/domain/models"
	preference_contracts_infra "github.com/junioralcant/api-stores-go/infra/contracts/preference_contracts"
)

type PreferenceCreateUseCase struct {
	Repo preference_contracts_infra.IPreferenceRepository
}

func NewPreferenceCreateUseCase(repo preference_contracts_infra.IPreferenceRepository) *PreferenceCreateUseCase {
	return &PreferenceCreateUseCase{Repo: repo}
}

func (p *PreferenceCreateUseCase) PreferenceCreate(preference models.Preference) *models.Preference {
	preferenceCreated := p.Repo.PreferenceCreateRepo(preference)

	return preferenceCreated
}
