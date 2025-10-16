package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"city-service/Controller"
	"city-service/Database"
)

func main() {
	if err := Database.InitDatabase(); err != nil {
    	log.Fatalf("Failed to initialize database: %v", err)
	}

	router := gin.Default()
	router.LoadHTMLFiles("View/index.html")

    router.GET("/", Controller.Index)
	router.GET("/cities", Controller.GetCities)
	router.GET("/city/:id", Controller.GetCityByID)
	router.POST("/city", Controller.CreateCity)
	router.PUT("/city/:id", Controller.UpdateCityByID)
	router.DELETE("/city/:id", Controller.DeleteCityByID)

	router.Run(":8080")
}
