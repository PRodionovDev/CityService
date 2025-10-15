package Database

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
	"city-service/Entity"
)

var DB *gorm.DB

func InitDatabase() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("cities.db"), &gorm.Config{})
	if err != nil {
    	return err
	}
	return DB.AutoMigrate(&Entity.City{})
}

func GetDB() *gorm.DB {
	return DB
}
