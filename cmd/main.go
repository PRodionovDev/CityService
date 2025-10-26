package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"city-service/internal/handler"
	"city-service/pkg/database"
	"city-service/pkg/middleware"
)

func main() {
	if err := database.InitDatabase(); err != nil {
    	log.Fatalf("Failed to initialize database: %v", err)
	}

	router := gin.Default()
	router.LoadHTMLFiles("internal/view/index.html")

    router.GET("/", handler.Index)

	router.Use(middleware.AuthMiddleware())
	router.GET("/cities", handler.GetCities)
	router.GET("/city/:id", handler.GetCityByID)
	router.POST("/city", handler.CreateCity)
	router.PUT("/city/:id", handler.UpdateCityByID)
	router.DELETE("/city/:id", handler.DeleteCityByID)
	router.GET("/regions", handler.GetRegions)
    router.GET("/region/:id", handler.GetRegionByID)
    router.POST("/region", handler.CreateRegion)
    router.PUT("/region/:id", handler.UpdateRegionByID)
    router.DELETE("/region/:id", handler.DeleteRegionByID)
    router.POST("/sync", handler.Sync)

	router.Run(":8080")
}
