package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	Dsn            string
	JWT_SECRET_KEY string
}

func LoadEnv() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Port:           os.Getenv("PORT"),
		Dsn:            os.Getenv("DSN"),
		JWT_SECRET_KEY: os.Getenv("JWT_SECRET_KEY"),
	}
}
