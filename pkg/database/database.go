package database

import (
    "fmt"
    "log"
    "os"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
	"city-service/internal/domain"
)

var DB *gorm.DB

func InitDatabase() error {
    dbHost := os.Getenv("DATABASE_HOST")
    dbName := os.Getenv("DATABASE_NAME")
    dbUser := os.Getenv("DATABASE_USER")
    dbPassword := os.Getenv("DATABASE_PASS")
    dbPort := os.Getenv("DATABASE_PORT")

    var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow", dbHost, dbUser, dbPassword, dbName, dbPort)
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
   		log.Fatalf("failed to connect database: %v", err)
   	}

    DB.AutoMigrate(&domain.Region{})
    return DB.AutoMigrate(&domain.City{})
}

func GetDB() *gorm.DB {
	return DB
}
