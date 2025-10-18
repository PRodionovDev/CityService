package Entity

import (
    "time"
)

type City struct {
	ID int	`json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	RegionID int `json:"region_id" "column:region_id"`
	IsCapital bool `json:"is_capital" "column:is_capital"`
	Type string `json:"type" gorm:"type:enum(\'Город\', \'Село\', \'пгт\', \'Деревня\')"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	TimeZone string `json:"time_zone"`
	Population int `json:"population"`
	Updated time.Time `gorm:"autoUpdateTime"`
}
