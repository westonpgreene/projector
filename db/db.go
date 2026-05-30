package db

import (
    "src/projector/models"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
    var err error
    DB, err = gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
    if err != nil {
        return err
    }
    return DB.AutoMigrate(&models.Order{}, &models.APIKey{})
}

