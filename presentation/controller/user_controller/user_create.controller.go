package user_controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/junioralcant/api-stores-go/domain/contracts/user_contracts"
	"github.com/junioralcant/api-stores-go/domain/models"
	"github.com/junioralcant/api-stores-go/presentation/helpers"
)

type UserCreateController struct {
	UseCase user_contracts.IUserCreate
}

func NewUserCreateController(useCase user_contracts.IUserCreate) *UserCreateController {
	return &UserCreateController{UseCase: useCase}
}

func (u *UserCreateController) Handle(ctx *gin.Context) {

	request := models.User{}

	ctx.BindJSON(&request)

	createUser := models.NewUser(request.Name, request.Email, request.Phone, request.CPF, request.Team, request.City, request.Street, request.PostalCode, request.Paid, request.ShirtSize, request.Distance, request.Sex, request.Age)

	fmt.Printf("createUser: %v\n", createUser)

	user, err := u.UseCase.UserCreate(*createUser)

	if err != nil {
		helpers.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.Header("Content-Type", "application/json")

	ctx.JSON(200, gin.H{
		"data": user,
	})
}
