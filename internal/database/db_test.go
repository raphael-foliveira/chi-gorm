package database

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/cfg"
)

func TestMain(m *testing.M) {
	err := cfg.LoadCfg("../../.env.test")
	if err != nil {
		panic(err)
	}
	m.Run()
}

func TestInitDb(t *testing.T) {
	t.Run("should initialize the database", func(t *testing.T) {
		Db = nil
		err := InitDb(cfg.DatabaseURL)
		if err != nil {
			t.Error(err)
		}
		if Db == nil {
			t.Error("Db not initialized")
		}
	})

	t.Run("should close the database", func(t *testing.T) {
		InitDb(cfg.DatabaseURL)
		err := CloseDb()
		if err != nil {
			t.Error(err)
		}
		if Db != nil {
			t.Error("Db not closed")
		}
	})
}
