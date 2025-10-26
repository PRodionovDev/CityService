package database

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
	"city-service/internal/domain"
)

var DB *gorm.DB

func InitDatabase() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("pkg/database/dump/cities.db"), &gorm.Config{})
	if err != nil {
    	return err
	}

    DB.AutoMigrate(&domain.Region{})
	return DB.AutoMigrate(&domain.City{})
}

func GetDB() *gorm.DB {
	return DB
}
