package service

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"city-service/internal/repository"
)

type CityService struct {
    repository repository.Cities
}

func NewCityService(repository repository.Cities) CityService {
    return CityService{repository: repository}
}

func (s *CityService) GetCities(c *gin.Context) {
    regionId := c.Query("region_id")
    name := c.Query("name")

	cities := s.repository.GetAllCities(name, regionId)
	c.JSON(http.StatusOK, cities)
}

func (s *CityService) GetCityByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid city ID"})
    	return
	}

	city := s.repository.GetCityByID(id)
	if city == nil {
    	c.JSON(http.StatusNotFound, gin.H{"error": "City not found"})
    	return
	}

	c.JSON(http.StatusOK, city)
}

func (s *CityService) CreateCity(c *gin.Context) {
	var input struct {
    	Name string         `json:"name" binding:"required"`
    	Slug string         `json:"slug" binding:"required"`
    	RegionID int        `json:"region_id" binding:"required"`
    	IsCapital bool      `json:"is_capital" binding:"required"`
    	Type string         `json:"type" binding:"required"`
    	Latitude float64    `json:"latitude" binding:"required"`
    	Longitude float64   `json:"longitude" binding:"required"`
    	TimeZone string     `json:"time_zone" binding:"required"`
    	Population int      `json:"population" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    	return
	}

	city := s.repository.CreateCity(input.Name, input.Slug, input.RegionID, input.IsCapital, input.Type, input.Latitude, input.Longitude, input.TimeZone, input.Population)
	c.JSON(http.StatusCreated, city)
}

func (s *CityService) UpdateCityByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid city ID"})
    	return
	}

	var input struct {
        	Name string         `json:"name" binding:"required"`
        	Slug string         `json:"slug" binding:"required"`
        	RegionID int        `json:"region_id" binding:"required"`
        	IsCapital bool      `json:"is_capital" binding:"required"`
        	Type string         `json:"type" binding:"required"`
        	Latitude float64    `json:"latitude" binding:"required"`
        	Longitude float64   `json:"longitude" binding:"required"`
        	TimeZone string     `json:"time_zone" binding:"required"`
        	Population int      `json:"population" binding:"required"`
    	}

	if err := c.ShouldBindJSON(&input); err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    	return
	}

	city := s.repository.UpdateCity(id, input.Name, input.Slug, input.RegionID, input.IsCapital, input.Type, input.Latitude, input.Longitude, input.TimeZone, input.Population)
	if city == nil {
    	c.JSON(http.StatusNotFound, gin.H{"error": "City not found"})
    	return
	}
	c.JSON(http.StatusOK, city)
}

func (s *CityService) DeleteCityByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid city ID"})
    	return
	}

	if success := s.repository.DeleteCityByID(id); !success {
    	c.JSON(http.StatusNotFound, gin.H{"error": "City not found"})
    	return
	}

	c.Status(http.StatusNoContent)
}