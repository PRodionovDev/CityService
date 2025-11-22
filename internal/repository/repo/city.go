package repo

import (
	"city-service/internal/domain"

	"gorm.io/gorm"
)

type CityRepository struct {
	db *gorm.DB
}

func NewCityRepository(db *gorm.DB) *CityRepository {
	return &CityRepository{db: db}
}

func (r *CityRepository) GetAllCities(name string, regionId string) []domain.City {
	var cities []domain.City
	db := r.db.Model(&domain.City{})

	if regionId != "" {
		db.Where("region_id = ?", regionId)
	}

	if name != "" {
		db.Where("name LIKE ?", "%"+name+"%")
	}

	db.Preload("Region").Find(&cities)
	return cities
}

func (r *CityRepository) GetCityByID(id int) *domain.City {
	var city domain.City
	if result := r.db.Preload("Region").First(&city, id); result.Error != nil {
		return nil
	}
	return &city
}

func (r *CityRepository) CreateCity(name string, slug string, regionId int, isCapital bool, cityType string, latitude float64, longitude float64, timeZone string, population int) domain.City {
	city := domain.City{
		Name:       name,
		Slug:       slug,
		RegionID:   regionId,
		IsCapital:  isCapital,
		Type:       cityType,
		Latitude:   latitude,
		Longitude:  longitude,
		TimeZone:   timeZone,
		Population: population,
	}

	r.db.Create(&city)
	return city
}

func (r *CityRepository) UpdateCity(id int, name string, slug string, regionId int, isCapital bool, cityType string, latitude float64, longitude float64, timeZone string, population int) *domain.City {
	var city domain.City
	if result := r.db.First(&city, id); result.Error != nil {
		return nil
	}
	city.Name = name
	city.Slug = slug
	city.RegionID = regionId
	city.IsCapital = isCapital
	city.Type = cityType
	city.Latitude = latitude
	city.Longitude = longitude
	city.TimeZone = timeZone
	city.Population = population

	r.db.Updates(&city)
	return &city
}

func (r *CityRepository) DeleteCityByID(id int) bool {
	if result := r.db.Delete(&domain.City{}, id); result.Error != nil {
		return false
	}
	return true
}
