package repository

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/cfg"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
)

func TestMain(m *testing.M) {
	err := cfg.LoadCfg("../../.env")
	if err != nil {
		panic(err)
	}
	err = database.InitDb(cfg.TestConfig.DatabaseURL)
	if err != nil {
		panic(err)
	}
	m.Run()
	err = database.CloseDb()
	if err != nil {
		panic(err)
	}
}
