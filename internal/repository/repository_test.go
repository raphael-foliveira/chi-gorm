package repository

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/cfg"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
)

func TestMain(m *testing.M) {
	err := cfg.LoadCfg("../../.env.test")
	if err != nil {
		panic(err)
	}
	config := cfg.GetCfg()
	_, err = database.GetDb(config.DatabaseURL)
	if err != nil {
		panic(err)
	}
	m.Run()
	err = database.CloseDb()
	if err != nil {
		panic(err)
	}
}
