package test

import (
	"city-service/internal/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

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
	db, mock := NewMockDB()
	repos := repository.NewRepositories(db)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO \"regions\" (\"name\",\"slug\",\"number\") VALUES ($1) RETURNING \"id\"").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	region := repos.Regions.CreateRegion("Moscow", "moscow", 99)
	if region.Name != "Moscow" {
		t.Fatalf("Failed to create region: %v", region)
	}
}

func TestUpdateRegionByID(t *testing.T) {
	db, mock := NewMockDB()
	repos := repository.NewRepositories(db)

	rows := sqlmock.NewRows([]string{"id", "name", "slug", "number"}).AddRow(1, "Mosow", "mosow", 99)
	mock.ExpectQuery("SELECT * FROM \"regions\" WHERE \"regions\".\"id\" = $1 ORDER BY \"regions\".\"id\" LIMIT $2").WillReturnRows(rows)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE \"users\" SET \"name\"=$1 WHERE \"id\" = $2").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	region := repos.Regions.UpdateRegion(1, "Moscow", "moscow", 99)
	if region.Name != "Moscow" {
		t.Fatalf("Failed to update region: %v", region)
	}
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
