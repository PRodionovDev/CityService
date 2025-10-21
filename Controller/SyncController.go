package Controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"encoding/csv"
    "fmt"
    "os"
    "city-service/Repository"
    "strconv"
)

func Sync(c *gin.Context) {
    ImportRegions()
    //ImportCities()
}

func ImportRegions() {
    regions := Import("Database/csv/regions.csv")

    for i := 0; i < len(regions); i++ {
        region := regions[i]
        name := region[0]
        slug := slug.Make(name)
        number, err := strconv.Atoi(region[1])
        if err != nil {
        	fmt.Println("Error converting string to int:", err)
        	return
        }
        Repository.CreateRegion(name, slug, number)
    }
}

func ImportCities() {
    cities := Import("Database/csv/cities.csv")
    fmt.Println(cities)
}

func Import(filename string) [][]string {
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