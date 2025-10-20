package Repository

import (
	"city-service/Entity"
	"city-service/Database"
)

/*func GetCityByName(name string) *Database.DB {
  return Database.DB.Where("name LIKE %?%", name)
}

func GetCityByRegionId(regionId string) *Database.DB {
  return Database.DB.Where("regionId = ?", regionId)
}*/

func GetAllCities(name string, regionId string) []Entity.City {
	var cities []Entity.City
	db := Database.DB.Model(&Entity.City{})

	if regionId != "" {
	    db.Where("region_id = ?", regionId)
	}

    if name != "" {
        db.Where("name LIKE ?", "%" + name + "%")
    }

	db.Find(&cities)
	return cities
}

func GetCityByID(id int) *Entity.City {
	var city Entity.City
	if result := Database.DB.First(&city, id); result.Error != nil {
    	return nil
	}
	return &city
}

func CreateCity(name string, slug string, regionId int, isCapital bool, cityType string, latitude float64, longitude float64, timeZone string, population int) Entity.City {
	city := Entity.City{
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

    Database.DB.Create(&city)
	return city
}

func UpdateCity(id int, name string, slug string, regionId int, isCapital bool, cityType string, latitude float64, longitude float64, timeZone string, population int) *Entity.City {
	var city Entity.City
    if result := Database.DB.First(&city, id); result.Error != nil {
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

	Database.DB.Updates(&city)
	return &city
}

func DeleteCityByID(id int) bool {
	if result := Database.DB.Delete(&Entity.City{}, id); result.Error != nil {
    	return false
	}
	return true
}
