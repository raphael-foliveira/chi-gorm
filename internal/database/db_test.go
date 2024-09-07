//go:build integration

package database_test

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/config"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
)

func TestMain(m *testing.M) {
	config.DatabaseURL = "postgres://postgres:postgres@localhost:5432/chi_gorm_test?sslmode=disable"
	m.Run()
}

func TestInitDb(t *testing.T) {
	t.Run("should retrieve a database instance", func(t *testing.T) {
		database.Start()
		if database.DB == nil {
			t.Error("Db not initialized")
		}
	})
}
