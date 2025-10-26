package repository

import (
	"city-service/internal/domain"
	"city-service/pkg/database"
)

func GetAllRegions(name string) []domain.Region {
	var regions []domain.Region
	db := database.DB.Model(&domain.Region{})

    if name != "" {
        db.Where("name LIKE ?", "%" + name + "%")
    }

    db.Find(&regions)
	return regions
}

func GetRegionByID(id int) *domain.Region {
	var region domain.Region
	if result := database.DB.First(&region, id); result.Error != nil {
    	return nil
	}
	return &region
}

func CreateRegion(name string, slug string, number int) domain.Region {
	region := domain.Region{
    	Name: name,
    	Slug: slug,
    	Number: number,
	}

    database.DB.Create(&region)
	return region
}

func UpdateRegion(id int, name string, slug string, number int) *domain.Region {
	var region domain.Region
    if result := database.DB.First(&region, id); result.Error != nil {
    	return nil
    }
	region.Name = name
	region.Slug = slug
	region.Number = number

	database.DB.Updates(&region)
	return &region
}

func DeleteRegionByID(id int) bool {
	if result := database.DB.Delete(&domain.Region{}, id); result.Error != nil {
    	return false
	}
	return true
}
