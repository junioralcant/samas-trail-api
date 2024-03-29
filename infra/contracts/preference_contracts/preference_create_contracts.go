package preference_contracts_infra

import "github.com/junioralcant/api-stores-go/domain/models"

type IPreferenceRepository interface {
	PreferenceCreateRepo(preference models.Preference) *models.Preference
}
