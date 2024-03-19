package router

import (
	"github.com/gin-gonic/gin"
	"github.com/junioralcant/api-stores-go/main/router/user_routes"
)

func Initialize() {
	apiPrefix := "/api/v1"

	r := gin.Default()

	user_routes.InitUserRoutes(r, apiPrefix)

	r.Run(":3333")
}
