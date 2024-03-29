package preference_routes

import (
	"github.com/gin-gonic/gin"
	"github.com/junioralcant/api-stores-go/main/factories/preference_factories"
)

func InitPreferenceRoutes(r *gin.Engine, apiPrefix string) {

	g := r.Group(apiPrefix)

	g.POST("/preference", preference_factories.PreferenceCreateControllerFactory().Handle)
}
