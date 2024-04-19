package user_repository

import (
	"fmt"

	"github.com/junioralcant/api-stores-go/domain/models"
	"github.com/junioralcant/api-stores-go/infra/config"
)

type UserCreateRepository struct {
}

func NewUserCreateRepository() *UserCreateRepository {
	return &UserCreateRepository{}
}

func (r *UserCreateRepository) UserCreateRepo(user models.User) (*models.User, error) {
	existingUser := &models.User{}
	if err := config.DB.Where("cpf = ?", user.CPF).First(existingUser).Error; err == nil {
		fmt.Println("User already exists")
		return nil, fmt.Errorf("CPF already exists")
	}

	fmt.Printf("existingUser: %v\n", existingUser)

	userCreate := models.User{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Phone:      user.Phone,
		CPF:        user.CPF,
		City:       user.City,
		Team:       user.Team,
		Street:     user.Street,
		PostalCode: user.PostalCode,
		ShirtSize:  user.ShirtSize,
		Paid:       user.Paid,
		Distance:   user.Distance,
		Sex:        user.Sex,
	}

	if err := config.DB.Create(&userCreate).Error; err != nil {
		config.Log.Errorf("error in create user: %v", err)
		return nil, err
	}

	return &userCreate, nil
}
