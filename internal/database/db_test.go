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
	t.Run("should retrieve a database instance", func(t *testing.T) {
		db, err := GetDb(cfg.Cfg.DatabaseURL)
		if err != nil {
			t.Error(err)
		}
		if db == nil {
			t.Error("Db not initialized")
		}
	})

	t.Run("should close the database", func(t *testing.T) {
		_, err := GetDb(cfg.Cfg.DatabaseURL)
		if err != nil {
			t.Error("could not initialize the database")
		}
		err = CloseDb()
		if err != nil {
			t.Error(err)
		}
		if GetInstance() != nil {
			t.Error("Db not closed")
		}
	})
}
