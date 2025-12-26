package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	RegisterAPIV1(router)
}

func RegisterAPIV1(router *gin.Engine) {
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			registerPublicHostRoutes(v1)
		}
	}
}
