package preference_repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/junioralcant/api-stores-go/domain/models"
	"github.com/junioralcant/api-stores-go/infra/config"
)

type PreferenceRepository struct {
}

func NewPreferenceRepository() *PreferenceRepository {
	return &PreferenceRepository{}
}

type Item struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	PictureURL  string  `json:"picture_url"`
	CategoryID  string  `json:"category_id"`
	Quantity    int     `json:"quantity"`
	CurrencyID  string  `json:"currency_id"`
	UnitPrice   float64 `json:"unit_price"`
}

type RequestBody struct {
	Items []Item `json:"items"`
}

func (r *PreferenceRepository) PreferenceCreateRepo(preference models.Preference) *models.Preference {

	url := "https://api.mercadopago.com/checkout/preferences"
	token := "APP_USR-3698788713698396-031409-407ff5620b593ab6635c3f4afb888985-186097166"

	// Adicionando um item ao array de items
	item := Item{
		Title:       preference.Title,
		Description: preference.Description,
		PictureURL:  "http://www.myapp.com/myimage.jpg",
		CategoryID:  preference.Category,
		Quantity:    1,
		CurrencyID:  preference.Currency,
		UnitPrice:   preference.Price,
	}

	request := RequestBody{
		Items: []Item{item},
	}

	// Converter a struct para JSON
	requestBody, err := json.Marshal(request)
	if err != nil {
		config.Log.Errorf("error in create preference: %v", err)
		return nil
	}

	// Criar um novo request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		config.Log.Errorf("error in create preference: %v", err)
		return nil
	}

	// Adicionar o cabeçalho com o token
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	fmt.Println("request:", req)

	// Enviar a requisição
	client := &http.Client{}
	resp, err := client.Do(req)
	fmt.Println("response:", resp.Body)

	if err != nil {
		config.Log.Errorf("error in create preference: %v", err)
		return nil
	}
	defer resp.Body.Close()

	var responsePreference struct {
		ID string `json:"id"`
	}

	err = json.NewDecoder(resp.Body).Decode(&responsePreference)
	if err != nil {
		config.Log.Errorf("error in create preference: %v", err)
		return nil
	}

	fmt.Println("responsePreference:", responsePreference)

	return &models.Preference{
		ID:          responsePreference.ID,
		CreatedAt:   preference.CreatedAt,
		UpdatedAt:   preference.UpdatedAt,
		Title:       preference.Title,
		Description: preference.Description,
		Category:    preference.Category,
		Price:       preference.Price,
		Currency:    preference.Currency,
	}
}
