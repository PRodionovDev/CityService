package test

import (
	"testing"
    "github.com/DATA-DOG/go-sqlmock"
    "city-service/internal/repository"
)

func TestFindAllCities(t *testing.T) {
    db, mock := NewMockDB()
    repos := repository.NewRepositories(db)

    rows := sqlmock.NewRows([]string{"id", "name", "slug", "region_id", "is_capital", "type", "latitude", "longitude", "time_zone", "population"}).
        AddRow(1, "Moscow", "moscow", 1, true, "Город", "55.751244", "37.618423", "Europe/Moscow", 15000000).
        AddRow(2, "Sankt-Peterburg", "spb", 1, false, "Город", "55.751244", "37.618423", "Europe/Moscow", 15000000)

    mock.ExpectQuery("SELECT * FROM \"cities\"").WillReturnRows(rows)

    cities := repos.Cities.GetAllCities("", "")
    if cities[0].Name != "Moscow" {
        t.Fatalf("Unexpected cities data retrieved: %v", cities)
    }
}

func TestFindCityByID(t *testing.T) {
    db, mock := NewMockDB()
    repos := repository.NewRepositories(db)

    rows := sqlmock.NewRows([]string{"id", "name", "slug", "region_id", "is_capital", "type", "latitude", "longitude", "time_zone", "population"}).
        AddRow(1, "Moscow", "moscow", 1, true, "Город", "55.751244", "37.618423", "Europe/Moscow", 15000000)
    regionRows := sqlmock.
    		NewRows([]string{"id", "name", "slug", "number"}).
    		AddRow(1, "moscow-oblast", "moscow-oblast", 99)

    mock.ExpectQuery("SELECT * FROM \"cities\" WHERE \"cities\".\"id\" = $1 ORDER BY \"cities\".\"id\" LIMIT $2").WillReturnRows(rows)
    mock.ExpectQuery("SELECT * FROM \"regions\" WHERE \"regions\".\"id\" = $1").WillReturnRows(regionRows)

    city := repos.Cities.GetCityByID(1)
    if city.Name != "Moscow" {
        t.Fatalf("Unexpected city data retrieved: %v", city)
    }
}

func TestUpdateCityByID(t *testing.T) {
    //
}

func TestDeleteCityByID(t *testing.T) {
    db, mock := NewMockDB()
    repos := repository.NewRepositories(db)

    mock.ExpectBegin()
    mock.ExpectExec("DELETE FROM \"cities\" WHERE \"cities\".\"id\" = $1").WillReturnResult(sqlmock.NewResult(1, 1))
    mock.ExpectCommit()

    isDeleted := repos.Cities.DeleteCityByID(1)
    if !isDeleted {
        t.Fatalf("Failed to delete city")
    }
}
