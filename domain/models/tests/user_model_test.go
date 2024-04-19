package tests

import (
	"testing"

	"github.com/junioralcant/api-stores-go/domain/models"
	"github.com/stretchr/testify/assert"
)

var userMocked = models.User{
	Name:       "jhon doe",
	Email:      "jhon@example.com",
	Phone:      "9999",
	CPF:        "111.111.111-11",
	Team:       "team",
	City:       "city",
	Street:     "street",
	PostalCode: "11111-111",
	Distance:   "10km",
	ShirtSize:  "M",
	Sex:        "F",
}

func TestUserModel_NewUser(t *testing.T) {
	paid := true
	user := models.NewUser(userMocked.Name, userMocked.Email, userMocked.Phone, userMocked.CPF, userMocked.Team, userMocked.City, userMocked.Street, userMocked.PostalCode, &paid, userMocked.ShirtSize, userMocked.Distance, userMocked.Sex)

	assert.EqualValues(t, user.Name, userMocked.Name)
	assert.EqualValues(t, user.Email, userMocked.Email)
	assert.EqualValues(t, user.Phone, userMocked.Phone)
	assert.EqualValues(t, user.CPF, userMocked.CPF)
	assert.EqualValues(t, user.Team, userMocked.Team)
	assert.EqualValues(t, user.City, userMocked.City)
	assert.EqualValues(t, user.PostalCode, userMocked.PostalCode)
	assert.EqualValues(t, user.Street, userMocked.Street)
	assert.EqualValues(t, user.Paid, &paid)
	assert.EqualValues(t, user.ShirtSize, userMocked.ShirtSize)
	assert.EqualValues(t, user.Distance, userMocked.Distance)
	assert.NotEmpty(t, user.ID)
	assert.EqualValues(t, user.Sex, userMocked.Sex)
}
