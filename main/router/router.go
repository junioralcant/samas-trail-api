package router

import (
	"github.com/gin-gonic/gin"
	"github.com/junioralcant/api-stores-go/main/router/preference_routes"
	"github.com/junioralcant/api-stores-go/main/router/user_routes"
	cors "github.com/rs/cors/wrapper/gin"
)

func Initialize() {
	apiPrefix := "/api/v1"

	r := gin.Default()

	r.Use(cors.AllowAll())

	user_routes.InitUserRoutes(r, apiPrefix)
	preference_routes.InitPreferenceRoutes(r, apiPrefix)

	r.Run(":3333")
}
