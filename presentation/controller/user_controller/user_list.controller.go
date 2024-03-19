package user_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/junioralcant/api-stores-go/domain/contracts/user_contracts"
)

type UserListController struct {
	UseCase user_contracts.IUserList
}

func NewUserListController(useCase user_contracts.IUserList) *UserListController {
	return &UserListController{UseCase: useCase}
}

func (u *UserListController) Handle(ctx *gin.Context) {
	users := u.UseCase.UserFindAll()
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, gin.H{
		"data": users,
	})
}
