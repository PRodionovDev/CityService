package repository

import (
	"city-service/internal/domain"
	"city-service/internal/repository/repo"

	"gorm.io/gorm"
)

type Cities interface {
	GetAllCities(name string, regionId string) []domain.City
	GetCityByID(id int) *domain.City
	CreateCity(name string, slug string, regionId int, isCapital bool, cityType string, latitude float64, longitude float64, timeZone string, population int) domain.City
	UpdateCity(id int, name string, slug string, regionId int, isCapital bool, cityType string, latitude float64, longitude float64, timeZone string, population int) *domain.City
	DeleteCityByID(id int) bool
}

type Regions interface {
	GetAllRegions(name string) []domain.Region
	GetRegionByID(id int) *domain.Region
	CreateRegion(name string, slug string, number int) domain.Region
	UpdateRegion(id int, name string, slug string, number int) *domain.Region
	DeleteRegionByID(id int) bool
}

type Repositories struct {
	Cities  Cities
	Regions Regions
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Cities:  repo.NewCityRepository(db),
		Regions: repo.NewRegionRepository(db),
	}
}
