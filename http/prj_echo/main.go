package main

import (
	"main/config"
	"main/internal/modules/item/model"
	"main/pkg/database"
	"main/pkg/routing"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.LoadConfig()
	db := database.NewGormDB(cfg)

	// Run migrations
	db.AutoMigrate(&model.Item{})

	e := echo.New()
	routing.RegisterRoutes(e, db)

	e.Logger.Fatal(e.Start(":8000"))
}
