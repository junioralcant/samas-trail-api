package tests

import (
	"testing"

	"github.com/junioralcant/api-stores-go/data/usecases/user_usecase"
	"github.com/junioralcant/api-stores-go/domain/models"
	"github.com/junioralcant/api-stores-go/infra/repositories/user_repository/mocks"
	"github.com/stretchr/testify/assert"
)

func makeSut() (*user_usecase.UserCreateUseCase, *mocks.UserCreateRepositorySpy) {
	repo := mocks.NewUserCreateRepositorySpy()
	useCase := user_usecase.NewUserCreateUseCase(repo)
	return useCase, repo
}

var paid = true

var userMocked = models.User{
	ID:         "19990",
	Name:       "test",
	Email:      "junior@test",
	City:       "city",
	Team:       "team",
	Street:     "street",
	PostalCode: "11111-111",
	Paid:       &paid,
	Phone:      "9999999999",
	CPF:        "48893993",
}

func TestUserCreateUseCase_correct_params(t *testing.T) {
	useCase, repoSpy := makeSut()

	useCase.UserCreate(userMocked)

	if repoSpy.Params != userMocked {
		t.Errorf("expected user %v, got %v", userMocked, repoSpy.Params)
	}
}

func TestUserCreateUseCase_correct_response(t *testing.T) {
	useCase, _ := makeSut()

	userCreated := useCase.UserCreate(userMocked)

	assert.EqualValues(t, userCreated.Name, userMocked.Name)
	assert.EqualValues(t, userCreated.Email, userMocked.Email)
	assert.EqualValues(t, userCreated.Phone, userMocked.Phone)
	assert.EqualValues(t, userCreated.CPF, userMocked.CPF)
	assert.EqualValues(t, userCreated.City, userMocked.City)
	assert.EqualValues(t, userCreated.Team, userMocked.Team)
	assert.EqualValues(t, userCreated.Street, userMocked.Street)
	assert.EqualValues(t, userCreated.PostalCode, userMocked.PostalCode)
	assert.EqualValues(t, userCreated.Paid, *userMocked.Paid)
	assert.EqualValues(t, userCreated.ID, userMocked.ID)

}
