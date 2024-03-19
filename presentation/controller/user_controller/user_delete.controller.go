package user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/junioralcant/api-stores-go/domain/contracts/user_contracts"
	"github.com/junioralcant/api-stores-go/presentation/helpers"
)

type UserDeleteController struct {
	UseCase user_contracts.IUserDelete
}

func NewUserDeleteController(useCase user_contracts.IUserDelete) *UserDeleteController {
	return &UserDeleteController{UseCase: useCase}
}

func (u *UserDeleteController) Handle(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		helpers.SendError(ctx, http.StatusBadRequest, helpers.ErrParamIsRequired("id", "string").Error())
		return
	}

	userDelete, err := u.UseCase.UserDelete(id)
	if err != nil {
		helpers.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SendSuccess(ctx, "user deleted", userDelete)
}
