package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"encoding/csv"
    "fmt"
    "os"
    "city-service/internal/repository"
    "strconv"
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
        slug := slug.Make(name)
        number, err := strconv.Atoi(region[1])
        if err != nil {
        	fmt.Println("Error converting string to int:", err)
        	return
        }
        s.regionRepo.CreateRegion(name, slug, number)
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
}

func importFile(filename string) [][]string {
    var data [][]string
    file, err := os.Open(filename)
    if err != nil {
    	panic(err)
    }
    defer file.Close()

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