package http

import (
    "github.com/gin-gonic/gin"
    "city-service/internal/service"
)

type Handler struct {
    CityService   service.CityService
    RegionService service.RegionService
    SiteService   service.SiteService
    SyncService   service.SyncService
}

func NewHandler(cityService service.CityService, regionService service.RegionService, siteService service.SiteService, syncService service.SyncService) *Handler {
    return &Handler{
        CityService:   cityService,
        RegionService: regionService,
        SiteService:   siteService,
        SyncService:   syncService,
    }
}

func (h *Handler) InitRoutes() *gin.Engine {
    router := gin.New()
    router.LoadHTMLFiles("internal/view/index.html")

    router.GET("/", h.SiteService.Index)

    router.Use(AuthMiddleware())
    router.GET("/cities", h.CityService.GetCities)
    router.GET("/city/:id", h.CityService.GetCityByID)
    router.POST("/city", h.CityService.CreateCity)
    router.PUT("/city/:id", h.CityService.UpdateCityByID)
    router.DELETE("/city/:id", h.CityService.DeleteCityByID)
    router.GET("/regions", h.RegionService.GetRegions)
    router.GET("/region/:id", h.RegionService.GetRegionByID)
    router.POST("/region", h.RegionService.CreateRegion)
    router.PUT("/region/:id", h.RegionService.UpdateRegionByID)
    router.DELETE("/region/:id", h.RegionService.DeleteRegionByID)
    router.POST("/sync", h.SyncService.Sync)

    return router
}