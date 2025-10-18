package Controller

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"city-service/Repository"
)

func GetRegions(c *gin.Context) {
	regions := Repository.GetAllRegions()
	c.JSON(http.StatusOK, regions)
}

func GetRegionByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid region ID"})
    	return
	}

	region := Repository.GetRegionByID(id)
	if region == nil {
    	c.JSON(http.StatusNotFound, gin.H{"error": "Region not found"})
    	return
	}

	c.JSON(http.StatusOK, region)
}

func CreateRegion(c *gin.Context) {
	var input struct {
    	Name string `json:"name" binding:"required"`
    	Slug string `json:"slug" binding:"required"`
    	Number int `json:"number" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    	return
	}

	region := Repository.CreateRegion(input.Name, input.Slug, input.Number)
	c.JSON(http.StatusCreated, region)
}

func UpdateRegionByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid region ID"})
    	return
	}

	var input struct {
    	Name string `json:"name" binding:"required"`
    	Slug string `json:"slug" binding:"required"`
    	Number int `json:"number" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    	return
	}

	region := Repository.UpdateRegion(id, input.Name, input.Slug, input.Number)
	if region == nil {
    	c.JSON(http.StatusNotFound, gin.H{"error": "Region not found"})
    	return
	}
	c.JSON(http.StatusOK, region)
}

func DeleteRegionByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid region ID"})
    	return
	}

	if success := Repository.DeleteRegionByID(id); !success {
    	c.JSON(http.StatusNotFound, gin.H{"error": "Region not found"})
    	return
	}

	c.Status(http.StatusNoContent)
}