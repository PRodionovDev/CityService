package Repository

import (
	"city-service/Entity"
	"city-service/Database"
)

func GetAllRegions() []Entity.Region {
	var regions []Entity.Region
	Database.DB.Find(&regions)
	return regions
}

func GetRegionByID(id int) *Entity.Region {
	var region Entity.Region
	if result := Database.DB.First(&region, id); result.Error != nil {
    	return nil
	}
	return &region
}

func CreateRegion(name string, slug string) Entity.Region {
	region := Entity.Region{
    	Name: name,
    	Slug: slug,
	}

    Database.DB.Create(&region)
	return region
}

func UpdateRegion(id int, name string, slug string) *Entity.Region {
	var region Entity.Region
    if result := Database.DB.First(&region, id); result.Error != nil {
    	return nil
    }
	region.Name = name
	region.Slug = slug

	Database.DB.Updates(&region)
	return &region
}

func DeleteRegionByID(id int) bool {
	if result := Database.DB.Delete(&Entity.Region{}, id); result.Error != nil {
    	return false
	}
	return true
}
