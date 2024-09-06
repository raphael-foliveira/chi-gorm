package database

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/config"
)

var testCfg *config.Cfg

func TestMain(m *testing.M) {
	testCfg = config.LoadCfg("../../.env.test")
	m.Run()
}

func TestInitDb(t *testing.T) {
	t.Run("should retrieve a database instance", func(t *testing.T) {
		DB := New(testCfg.DatabaseURL)
		if DB == nil {
			t.Error("Db not initialized")
		}
	})
}
