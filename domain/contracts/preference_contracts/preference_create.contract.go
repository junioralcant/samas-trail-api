package preference_contracts

import "github.com/junioralcant/api-stores-go/domain/models"

type IPreferenceCreate interface {
	PreferenceCreate(preference models.Preference) *models.Preference
}
