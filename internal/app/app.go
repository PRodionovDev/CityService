package app

import (
	"log"
	"city-service/internal/delivery/http"
	"city-service/internal/server"
	"city-service/internal/repository"
	"city-service/internal/service"
	"city-service/pkg/database"
)

func Run() {
	if err := database.InitDatabase(); err != nil {
    	log.Fatalf("Failed to initialize database: %v", err)
	}

    repos := repository.NewRepositories(database.DB)
    handlers := http.NewHandler(
        service.NewCityService(repos.Cities),
        service.NewRegionService(repos.Regions),
        service.NewSiteService(),
        service.NewSyncService(repos.Regions, repos.Cities),
    )

	srv := new(server.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
	    log.Fatalf("Failed to run http server: %e", err.Error())
	}
}
