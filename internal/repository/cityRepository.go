package repository

import (
	"city-service/internal/domain"
	"city-service/pkg/database"
)

func GetAllCities(name string, regionId string) []domain.City {
	var cities []domain.City
	db := database.DB.Model(&domain.City{})

	if regionId != "" {
	    db.Where("region_id = ?", regionId)
	}

    if name != "" {
        db.Where("name LIKE ?", "%" + name + "%")
    }

	db.Find(&cities)
	return cities
}

func GetCityByID(id int) *domain.City {
	var city domain.City
	if result := database.DB.First(&city, id); result.Error != nil {
    	return nil
	}
	return &city
}

func CreateCity(name string, slug string, regionId int, isCapital bool, cityType string, latitude float64, longitude float64, timeZone string, population int) domain.City {
	city := domain.City{
    	Name: name,
    	Slug: slug,
    	RegionID: regionId,
    	IsCapital: isCapital,
    	Type: cityType,
    	Latitude: latitude,
    	Longitude: longitude,
    	TimeZone: timeZone,
    	Population: population,
	}

    database.DB.Create(&city)
	return city
}

func UpdateCity(id int, name string, slug string, regionId int, isCapital bool, cityType string, latitude float64, longitude float64, timeZone string, population int) *domain.City {
	var city domain.City
    if result := database.DB.First(&city, id); result.Error != nil {
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

	database.DB.Updates(&city)
	return &city
}

func DeleteCityByID(id int) bool {
	if result := database.DB.Delete(&domain.City{}, id); result.Error != nil {
    	return false
	}
	return true
}
