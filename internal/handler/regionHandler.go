package handler

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"city-service/internal/repository"
)

func GetRegions(c *gin.Context) {
    name := c.Query("name")

	regions := repository.GetAllRegions(name)
	c.JSON(http.StatusOK, regions)
}

func GetRegionByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid region ID"})
    	return
	}

	region := repository.GetRegionByID(id)
	if region == nil {
    	c.JSON(http.StatusNotFound, gin.H{"error": "Region not found"})
    	return
	}

	c.JSON(http.StatusOK, region)
}

func CreateRegion(c *gin.Context) {
	var input struct {
    	Name string     `json:"name" binding:"required"`
    	Slug string     `json:"slug" binding:"required"`
    	Number int      `json:"number" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    	return
	}

	region := repository.CreateRegion(input.Name, input.Slug, input.Number)
	c.JSON(http.StatusCreated, region)
}

func UpdateRegionByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid region ID"})
    	return
	}

	var input struct {
    	Name string     `json:"name" binding:"required"`
    	Slug string     `json:"slug" binding:"required"`
    	Number int      `json:"number" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    	return
	}

	region := repository.UpdateRegion(id, input.Name, input.Slug, input.Number)
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

	if success := repository.DeleteRegionByID(id); !success {
    	c.JSON(http.StatusNotFound, gin.H{"error": "Region not found"})
    	return
	}

	c.Status(http.StatusNoContent)
}