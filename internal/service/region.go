package service

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"city-service/internal/repository"
)

type RegionService struct {
    repository repository.Regions
}

func NewRegionService(repository repository.Regions) RegionService {
    return RegionService{repository: repository}
}

func (s *RegionService) GetRegions(c *gin.Context) {
    name := c.Query("name")

	regions := s.repository.GetAllRegions(name)
	c.JSON(http.StatusOK, regions)
}

func (s *RegionService) GetRegionByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid region ID"})
    	return
	}

	region := s.repository.GetRegionByID(id)
	if region == nil {
    	c.JSON(http.StatusNotFound, gin.H{"error": "Region not found"})
    	return
	}

	c.JSON(http.StatusOK, region)
}

func (s *RegionService) CreateRegion(c *gin.Context) {
	var input struct {
    	Name string     `json:"name" binding:"required"`
    	Slug string     `json:"slug" binding:"required"`
    	Number int      `json:"number" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    	return
	}

	region := s.repository.CreateRegion(input.Name, input.Slug, input.Number)
	c.JSON(http.StatusCreated, region)
}

func (s *RegionService) UpdateRegionByID(c *gin.Context) {
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

	region := s.repository.UpdateRegion(id, input.Name, input.Slug, input.Number)
	if region == nil {
    	c.JSON(http.StatusNotFound, gin.H{"error": "Region not found"})
    	return
	}
	c.JSON(http.StatusOK, region)
}

func (s *RegionService) DeleteRegionByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid region ID"})
    	return
	}

	if success := s.repository.DeleteRegionByID(id); !success {
    	c.JSON(http.StatusNotFound, gin.H{"error": "Region not found"})
    	return
	}

	c.Status(http.StatusNoContent)
}