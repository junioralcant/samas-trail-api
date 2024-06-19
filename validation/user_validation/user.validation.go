package user_validation

import (
	"fmt"

	"github.com/junioralcant/api-stores-go/domain/models"
)

type UpdateUser struct {
	User models.User
}

func (r *UpdateUser) Validate() error {
	if r.User.IdPreferences != "" || r.User.CPF != "" || r.User.Email != "" || r.User.Name != "" || r.User.City != "" || r.User.Street != "" || r.User.PostalCode != "" || r.User.Phone != "" || r.User.Team != "" || r.User.Paid != nil || r.User.ShirtSize != "" || r.User.Distance != "" || r.User.Sex != "" || r.User.Age != "" || !r.User.CreatedAt.IsZero() {
		return nil
	}

	return fmt.Errorf("at least one valid field must be provided")
}
