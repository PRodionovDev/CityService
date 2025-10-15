package Repository

import (
	"city-service/Entity"
	"city-service/Database"
)

func GetAllCities() []Entity.City {
	var cities []Entity.City
	Database.DB.Find(&cities)
	return cities
}

func GetCityByID(id int) *Entity.City {
	var city Entity.City
	if result := Database.DB.First(&city, id); result.Error != nil {
    	return nil
	}
	return &city
}

func CreateCity(name string, slug string) Entity.City {
	city := Entity.City{
    	Name: name,
    	Slug: slug,
	}

    Database.DB.Create(&city)
	return city
}

func UpdateCity(id int, name string, slug string) *Entity.City {
	var city Entity.City
    if result := Database.DB.First(&city, id); result.Error != nil {
    	return nil
    }
	city.Name = name
	city.Slug = slug

	Database.DB.Updates(&city)
	return &city
}

func DeleteCityByID(id int) bool {
	if result := Database.DB.Delete(&Entity.City{}, id); result.Error != nil {
    	return false
	}
	return true
}
