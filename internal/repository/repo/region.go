package repo

import (
	"city-service/internal/domain"

	"gorm.io/gorm"
)

type RegionRepository struct {
	db *gorm.DB
}

func NewRegionRepository(db *gorm.DB) *RegionRepository {
	return &RegionRepository{db: db}
}

func (r *RegionRepository) GetAllRegions(name string) []domain.Region {
	var regions []domain.Region
	db := r.db.Model(&domain.Region{})

	if name != "" {
		db.Where("name LIKE ?", "%"+name+"%")
	}

	db.Find(&regions)
	return regions
}

func (r *RegionRepository) GetRegionByID(id int) *domain.Region {
	var region domain.Region
	if result := r.db.First(&region, id); result.Error != nil {
		return nil
	}
	return &region
}

func (r *RegionRepository) CreateRegion(name string, slug string, number int) domain.Region {
	region := domain.Region{
		Name:   name,
		Slug:   slug,
		Number: number,
	}

	r.db.Create(&region)
	return region
}

func (r *RegionRepository) UpdateRegion(id int, name string, slug string, number int) *domain.Region {
	var region domain.Region
	if result := r.db.First(&region, id); result.Error != nil {
		return nil
	}
	region.Name = name
	region.Slug = slug
	region.Number = number

	r.db.Updates(&region)
	return &region
}

func (r *RegionRepository) DeleteRegionByID(id int) bool {
	if result := r.db.Delete(&domain.Region{}, id); result.Error != nil {
		return false
	}
	return true
}
