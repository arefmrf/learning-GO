package routes

import (
	"trip/internal/services/public_host"

	"github.com/gin-gonic/gin"
)

func registerPublicHostRoutes(r *gin.RouterGroup) {
	hosts := r.Group("public/hosts")
	{
		hosts.GET("/general", GetPublicHosts)
	}
}

func GetPublicHosts(c *gin.Context) {
	items, err := public_host.GetPublicHosts(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": "internal_error"})
		return
	}

	c.JSON(200, items)
}
