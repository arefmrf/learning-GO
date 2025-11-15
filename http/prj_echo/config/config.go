package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl      string
	DBPort     string
	DBPassword string
	DBName     string
	DBUser     string
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	return &Config{
		DBUrl:      os.Getenv("DATABASE_URL"),
		DBPort:     os.Getenv("DATABASE_PORT"),
		DBPassword: os.Getenv("DATABASE_PASSWORD"),
		DBName:     os.Getenv("DATABASE_PASSWORD"),
		DBUser:     os.Getenv("DATABASE_PASSWORD"),
	}
}
