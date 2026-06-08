package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *Config) *gorm.DB {
	dsn := cfg.Dsn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	return db
}
