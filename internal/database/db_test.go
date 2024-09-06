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
		Initialize(testCfg.DatabaseURL)
		if DB == nil {
			t.Error("Db not initialized")
		}
	})

	t.Run("should close the database", func(t *testing.T) {
		Initialize(testCfg.DatabaseURL)
		err := Close()
		if err != nil {
			t.Error(err)
		}
		if DB != nil {
			t.Error("Db not closed")
		}
	})
}
