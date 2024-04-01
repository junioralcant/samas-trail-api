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

type URLS struct {
	Success string `json:"success"`
	Failure string `json:"failure"`
	Pending string `json:"pending"`
}

type RequestBody struct {
	Items      []Item `json:"items"`
	BackURLS   URLS   `json:"back_urls"`
	AutoReturn string `json:"auto_return"`
}

func (r *PreferenceRepository) PreferenceCreateRepo(preference models.Preference) *models.Preference {

	url := "https://api.mercadopago.com/checkout/preferences"
	token := config.GetEnvVariable("TOKEN_MERCADO_PAGO")

	// Adicionando um item ao array de items
	item := Item{
		Title:       preference.Title,
		Description: preference.Description,
		PictureURL:  "",
		CategoryID:  preference.Category,
		Quantity:    1,
		CurrencyID:  preference.Currency,
		UnitPrice:   preference.Price,
	}

	request := RequestBody{
		Items: []Item{item},
		BackURLS: URLS{
			Success: "http://localhost:3000/cadastro",
			Failure: "http://localhost:3000/cadastro",
			Pending: "http://localhost:3000/cadastro",
		},
		AutoReturn: "all",
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
