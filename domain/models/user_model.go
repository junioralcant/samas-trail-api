package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            string    `json:"id" gorm:"type:uuid; primarykey;"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	Name          string    `json:"name"`
	Phone         string    `json:"phone"`
	CPF           string    `json:"cpf"`
	Email         string    `json:"email"`
	Team          string    `json:"team"`
	City          string    `json:"city"`
	Street        string    `json:"street"`
	PostalCode    string    `json:"postalCode"`
	Paid          *bool     `json:"paid" gorm:"default:false"`
	ShirtSize     string    `json:"shirtSize"`
	Distance      string    `json:"distance"`
	IdPreferences string    `json:"idPreferences"`
}

func NewUser(name string, email string, phone string, cpf string, team string, city string, street string, postalCode string, paid *bool, shirtSize string, distance string) *User {
	user := User{
		ID:            uuid.New().String(),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		Name:          name,
		Email:         email,
		Phone:         phone,
		CPF:           cpf,
		City:          city,
		Team:          team,
		Street:        street,
		PostalCode:    postalCode,
		ShirtSize:     shirtSize,
		Paid:          paid,
		Distance:      distance,
		IdPreferences: "",
	}

	return &user
}
