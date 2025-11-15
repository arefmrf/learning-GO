package main

import (
	"main/config"
	"main/internal/modules/item/handler"
	"main/internal/modules/item/model"
	"main/internal/modules/item/repository"
	"main/internal/modules/item/service"
	"main/pkg/database"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.LoadConfig()
	db := database.NewGormDB(cfg)

	// Run migrations
	db.AutoMigrate(&model.Item{})

	itemRepo := repository.NewItemRepository(db)
	itemSvc := service.NewItemService(itemRepo)
	itemHandler := handler.NewItemHandler(itemSvc)

	e := echo.New()

	// CRUD routes
	e.POST("/items", itemHandler.Create)
	e.GET("/items", itemHandler.List)
	e.GET("/items/:id", itemHandler.GetByID)
	e.PUT("/items/:id", itemHandler.Update)
	e.DELETE("/items/:id", itemHandler.Delete)

	e.Logger.Fatal(e.Start(":8000"))
}
