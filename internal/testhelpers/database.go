package testhelpers

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/database"
)

func StartDB(t *testing.T) {
	if err := database.Start(); err != nil {
		t.Log("Error starting test database")
		t.Fatal(err)
	}
}
