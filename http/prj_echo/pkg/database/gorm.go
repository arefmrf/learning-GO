package database

import (
	"log"
	"main/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB(cfg *config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.DBUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	return db
}
