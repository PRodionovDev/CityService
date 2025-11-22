package service

import (
	"city-service/internal/repository"
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

type SyncService struct {
	regionRepo repository.Regions
	cityRepo   repository.Cities
}

func NewSyncService(regionRepo repository.Regions, cityRepo repository.Cities) SyncService {
	return SyncService{regionRepo: regionRepo, cityRepo: cityRepo}
}

func (s *SyncService) Sync(c *gin.Context) {
	regions := importFile("pkg/database/csv/regions.csv")

	for i := 0; i < len(regions); i++ {
		region := regions[i]
		name := region[0]
		slugName := slug.Make(name)
		number, err := strconv.Atoi(region[1])
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return
		}
		s.regionRepo.CreateRegion(name, slugName, number)
	}

	cities := importFile("pkg/database/csv/cities.csv")

	for i := 0; i < len(cities); i++ {
		city := cities[i]
		region := s.regionRepo.GetAllRegions(city[4])
		lat, err := strconv.ParseFloat(city[2], 64)
		if err != nil {
			fmt.Println("Error converting string to float64:", err)
			return
		}
		long, err := strconv.ParseFloat(city[2], 64)
		if err != nil {
			fmt.Println("Error converting string to float64:", err)
			return
		}

		// Not enough data, so:
		// isCapital - default false
		// cityType - default "Город"
		// timezone - default ""
		// timezone - default 0
		s.cityRepo.CreateCity(city[0], city[1], region[0].ID, false, "Город", lat, long, "", 0)
	}
	c.JSON(http.StatusOK, gin.H{})
}

func importFile(filename string) [][]string {
	var data [][]string
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	reader := csv.NewReader(file)

	for {
		record, e := reader.Read()
		if e != nil {
			fmt.Println(e)
			break
		}
		data = append(data, record)
	}

	return data
}
