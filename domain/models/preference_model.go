package models

import (
	"time"
)

type Preference struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Currency    string    `json:"currency"`
	Price       float64   `json:"price"`
}

func NewPreference(title string, description string, category string, currency string, price float64) *Preference {
	preference := Preference{
		ID:          "",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Title:       title,
		Description: description,
		Category:    category,
		Currency:    currency,
		Price:       price,
	}

	return &preference
}
