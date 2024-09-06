package database

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/config"
)

func TestMain(m *testing.M) {
	config.Initialize("../../.env.test")
	m.Run()
}

func TestInitDb(t *testing.T) {
	t.Run("should retrieve a database instance", func(t *testing.T) {
		Initialize(config.DatabaseURL)
		if DB == nil {
			t.Error("Db not initialized")
		}
	})
}
