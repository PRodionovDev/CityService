package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"city-service/Controller"
	"city-service/Database"
	"city-service/Middleware"
)

func main() {
	if err := Database.InitDatabase(); err != nil {
    	log.Fatalf("Failed to initialize database: %v", err)
	}

	router := gin.Default()
	router.LoadHTMLFiles("View/index.html")

    router.GET("/", Controller.Index)

	router.Use(Middleware.AuthMiddleware())
	router.GET("/cities", Controller.GetCities)
	router.GET("/city/:id", Controller.GetCityByID)
	router.POST("/city", Controller.CreateCity)
	router.PUT("/city/:id", Controller.UpdateCityByID)
	router.DELETE("/city/:id", Controller.DeleteCityByID)
	router.GET("/regions", Controller.GetRegions)
    router.GET("/region/:id", Controller.GetRegionByID)
    router.POST("/region", Controller.CreateRegion)
    router.PUT("/region/:id", Controller.UpdateRegionByID)
    router.DELETE("/region/:id", Controller.DeleteRegionByID)
    router.POST("/sync", Controller.Sync)

	router.Run(":8080")
}
