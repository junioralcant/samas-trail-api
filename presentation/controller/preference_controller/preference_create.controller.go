package preference_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/junioralcant/api-stores-go/domain/contracts/preference_contracts"
	"github.com/junioralcant/api-stores-go/domain/models"
)

type PreferenceCreateController struct {
	UseCase preference_contracts.IPreferenceCreate
}

func NewPreferenceCreateController(useCase preference_contracts.IPreferenceCreate) *PreferenceCreateController {
	return &PreferenceCreateController{UseCase: useCase}
}

func (p *PreferenceCreateController) Handle(ctx *gin.Context) {

	request := models.Preference{}

	ctx.BindJSON(&request)

	createP := models.NewPreference(request.Title, request.Description, request.Category, request.Currency, request.Price)

	preference := p.UseCase.PreferenceCreate(*createP)

	ctx.Header("Content-Type", "application/json")

	ctx.JSON(200, gin.H{
		"data": preference,
	})
}
