package database_test

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/config"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
)

func TestMain(m *testing.M) {
	config.Load("../../.env.test")
	m.Run()
}

func TestInitDb(t *testing.T) {
	t.Run("should retrieve a database instance", func(t *testing.T) {
		database.Start(config.DatabaseURL)
		if database.DB == nil {
			t.Error("Db not initialized")
		}
	})
}
