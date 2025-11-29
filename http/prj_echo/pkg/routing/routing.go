package routing

import (
	"main/internal/routes"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	routes.ModuleRoutes(e, db)
}
