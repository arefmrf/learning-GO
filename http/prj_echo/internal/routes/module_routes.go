package routes

import (
	"main/internal/modules/item/handler"
	"main/internal/modules/item/repository"
	"main/internal/modules/item/service"
	userController "main/internal/modules/user/handler"
	"main/pkg/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func ModuleRoutes(e *echo.Echo, db *gorm.DB) {
	itemRepo := repository.NewItemRepository(db)
	itemSvc := service.NewItemService(itemRepo)
	itemHandler := handler.NewItemHandler(itemSvc)

	//e.Use(middleware.Logger())
	//e.Use(middlewares.CustomMiddleware)

	api := e.Group("/api", middleware.Logger())
	// CRUD routes
	api.POST("/items", itemHandler.Create, middlewares.CustomMiddleware)
	api.GET("/items", itemHandler.List)
	api.GET("/items/:id", itemHandler.GetByID)
	api.PUT("/items/:id", itemHandler.Update)
	api.DELETE("/items/:id", itemHandler.Delete)

	user := userController.NewController()
	api.POST("/items", user.Register)

}
