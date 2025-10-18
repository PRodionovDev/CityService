package Controller

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"city-service/Repository"
)

func GetCities(c *gin.Context) {
	cities := Repository.GetAllCities()
	c.JSON(http.StatusOK, cities)
}

func GetCityByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid city ID"})
    	return
	}

	city := Repository.GetCityByID(id)
	if city == nil {
    	c.JSON(http.StatusNotFound, gin.H{"error": "City not found"})
    	return
	}

	c.JSON(http.StatusOK, city)
}

func CreateCity(c *gin.Context) {
	var input struct {
    	Name string `json:"name" binding:"required"`
    	Slug string `json:"slug" binding:"required"`
    	RegionID int `json:"region_id" binding:"required"`
    	IsCapital bool `json:"is_capital" binding:"required"`
    	Type string `json:"type" binding:"required"`
    	Latitude float64 `json:"latitude" binding:"required"`
    	Longitude float64 `json:"longitude" binding:"required"`
    	TimeZone string `json:"time_zone" binding:"required"`
    	Population int `json:"population" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    	return
	}

	city := Repository.CreateCity(input.Name, input.Slug, input.RegionID, input.IsCapital, input.Type, input.Latitude, input.Longitude, input.TimeZone, input.Population)
	c.JSON(http.StatusCreated, city)
}

func UpdateCityByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid city ID"})
    	return
	}

	var input struct {
        	Name string `json:"name" binding:"required"`
        	Slug string `json:"slug" binding:"required"`
        	RegionID int `json:"region_id" binding:"required"`
        	IsCapital bool `json:"is_capital" binding:"required"`
        	Type string `json:"type" binding:"required"`
        	Latitude float64 `json:"latitude" binding:"required"`
        	Longitude float64 `json:"longitude" binding:"required"`
        	TimeZone string `json:"time_zone" binding:"required"`
        	Population int `json:"population" binding:"required"`
    	}

	if err := c.ShouldBindJSON(&input); err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    	return
	}

	city := Repository.UpdateCity(id, input.Name, input.Slug, input.RegionID, input.IsCapital, input.Type, input.Latitude, input.Longitude, input.TimeZone, input.Population)
	if city == nil {
    	c.JSON(http.StatusNotFound, gin.H{"error": "City not found"})
    	return
	}
	c.JSON(http.StatusOK, city)
}

func DeleteCityByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid city ID"})
    	return
	}

	if success := Repository.DeleteCityByID(id); !success {
    	c.JSON(http.StatusNotFound, gin.H{"error": "City not found"})
    	return
	}

	c.Status(http.StatusNoContent)
}