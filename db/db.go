package db

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"src/projector/models"
)

var DB *gorm.DB

func Init() error {
	if err := os.MkdirAll("data", 0755); err != nil {
		return err
	}
	var err error
	DB, err = gorm.Open(sqlite.Open("data/orders.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	return DB.AutoMigrate(&models.Order{}, &models.APIKey{})
}
