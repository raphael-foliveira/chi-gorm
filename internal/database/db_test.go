package database

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/config"
)

func TestMain(m *testing.M) {
	config.LoadCfg("../../.env.test")
	m.Run()
}

func TestInitDb(t *testing.T) {
	t.Run("should retrieve a database instance", func(t *testing.T) {
		db := Db()
		if db == nil {
			t.Error("Db not initialized")
		}
	})

	t.Run("should close the database", func(t *testing.T) {
		Db()
		err := Close()
		if err != nil {
			t.Error(err)
		}
		if instance != nil {
			t.Error("Db not closed")
		}
	})
}
