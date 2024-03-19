package user_routes

import (
	"github.com/gin-gonic/gin"
	"github.com/junioralcant/api-stores-go/main/factories/user_factories"
)

func InitUserRoutes(r *gin.Engine, apiPrefix string) {

	g := r.Group(apiPrefix)

	g.GET("/users", user_factories.UserListAllControllerFactory().Handle)

	g.POST("/user", user_factories.UserCreateControllerFactory().Handle)

	g.PUT("/user", user_factories.UserUpdateControllerFactory().Handle)

	g.DELETE("/user", user_factories.UserDeleteControllerFactory().Handle)
}
