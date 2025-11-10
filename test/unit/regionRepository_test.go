package test

import (
	"testing"
	"log"
	"gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/DATA-DOG/go-sqlmock"
    "city-service/internal/repository"
)

func NewMockDB() (*gorm.DB, sqlmock.Sqlmock) {
    db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
    if err != nil {
        log.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
    }

    gormDB, err := gorm.Open(postgres.New(postgres.Config{
        Conn: db,
    }), &gorm.Config{})

    if err != nil {
        log.Fatalf("An error '%s' was not expected when opening gorm database", err)
    }

    return gormDB, mock
}

func TestFindAllRegions(t *testing.T) {
    db, mock := NewMockDB()
    repos := repository.NewRepositories(db)

    rows := sqlmock.NewRows([]string{"id", "name", "slug", "number"}).
        AddRow(1, "Moscow", "moscow", 99).
        AddRow(1, "Sankt-Peterburg", "spb", 77)

    mock.ExpectQuery("SELECT * FROM \"regions\"").WillReturnRows(rows)

    regions := repos.Regions.GetAllRegions("")
    if regions[0].Name != "Moscow" {
        t.Fatalf("Unexpected regions data retrieved: %v", regions)
    }
}

func TestFindRegionByID(t *testing.T) {
    db, mock := NewMockDB()
    repos := repository.NewRepositories(db)

    rows := sqlmock.NewRows([]string{"id", "name", "slug", "number"}).
        AddRow(1, "Moscow", "moscow", 99)

    mock.ExpectQuery("SELECT * FROM \"regions\" WHERE \"regions\".\"id\" = $1 ORDER BY \"regions\".\"id\" LIMIT $2").WillReturnRows(rows)

    region := repos.Regions.GetRegionByID(1)
    if region.Name != "Moscow" {
        t.Fatalf("Unexpected region data retrieved: %v", region)
    }
}

func TestCreateRegion(t *testing.T) {
    /*db, mock := NewMockDB()
    repos := repository.NewRepositories(db)

    rows := sqlmock.NewRows([]string{"id", "name", "slug", "number"}).
        AddRow(1, "Moscow", "moscow", 99)

    mock.ExpectQuery("SELECT * FROM \"regions\" WHERE \"regions\".\"id\" = $1 ORDER BY \"regions\".\"id\" LIMIT $2").WillReturnRows(rows)

    region := repos.Regions.GetRegionByID(1)
    if region.Name != "Moscow" {
        t.Fatalf("Unexpected region data retrieved: %v", region)
    }*/
}

func TestUpdateRegionByID(t *testing.T) {
    /*db, mock := NewMockDB()
    repos := repository.NewRepositories(db)

    rows := sqlmock.NewRows([]string{"id", "name", "slug", "number"}).
        AddRow(1, "Moscow", "moscow", 99)

    mock.ExpectQuery("SELECT * FROM \"regions\" WHERE \"regions\".\"id\" = $1 ORDER BY \"regions\".\"id\" LIMIT $2").WillReturnRows(rows)

    region := repos.Regions.GetRegionByID(1)
    if region.Name != "Moscow" {
        t.Fatalf("Unexpected region data retrieved: %v", region)
    }*/
}

func TestDeleteRegionByID(t *testing.T) {
    db, mock := NewMockDB()
    repos := repository.NewRepositories(db)

    mock.ExpectBegin()
    mock.ExpectExec("DELETE FROM \"regions\" WHERE \"regions\".\"id\" = $1").WillReturnResult(sqlmock.NewResult(1, 1))
    mock.ExpectCommit()

    isDeleted := repos.Regions.DeleteRegionByID(1)
    if !isDeleted {
        t.Fatalf("Failed to delete region")
    }
}
