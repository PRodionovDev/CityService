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
    	//RegionID int `json:"number" binding:"required"`
    	//IsCapital int `json:"number" binding:"required"`
    	//Type int `json:"number" binding:"required"`
    	//Latitude int `json:"number" binding:"required"`
    	//Longitude int `json:"number" binding:"required"`
    	//TimeZone int `json:"number" binding:"required"`
    	//Population int `json:"number" binding:"required"`
    	//Updated int `json:"number" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    	return
	}

	city := Repository.CreateCity(input.Name, input.Slug)
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
	}

	if err := c.ShouldBindJSON(&input); err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    	return
	}

	city := Repository.UpdateCity(id, input.Name, input.Slug)
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