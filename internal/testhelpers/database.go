package testhelpers

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/config"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
)

var testDbUrl = "postgres://postgres:postgres@localhost:5432/chi_gorm_test?sslmode=disable"

func StartDB(t *testing.T) {
	config.DatabaseURL = testDbUrl
	if err := database.Start(); err != nil {
		t.Log("Error starting test database")
		t.Fatal(err)
	}
}
