package database

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
	"city-service/internal/domain"
)

var DB *gorm.DB

func InitDatabase(environment string) error {
    name := "cities.db"
    if environment == "test" {
        name = "cities_test.db"
    }

	var err error
	DB, err = gorm.Open(sqlite.Open("pkg/database/dump/" + name), &gorm.Config{})
	if err != nil {
    	return err
	}

    DB.AutoMigrate(&domain.Region{})
	return DB.AutoMigrate(&domain.City{})
}

func GetDB() *gorm.DB {
	return DB
}
