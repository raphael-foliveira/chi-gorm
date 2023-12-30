package database

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/cfg"
)

func TestInitDb(t *testing.T) {
	cfg.LoadEnv("../../.env")
	t.Run("should initialize the database", func(t *testing.T) {
		Db = nil
		err := InitDb(cfg.TestConfig.DatabaseURL)
		if err != nil {
			t.Error(err)
		}
		if Db == nil {
			t.Error("Db not initialized")
		}
	})
}
