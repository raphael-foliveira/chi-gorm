package repository

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/cfg"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
)

func TestMain(m *testing.M) {
	cfg.LoadCfg("../../.env.test")
	database.Db()
	m.Run()
	err := database.CloseDb()
	if err != nil {
		panic(err)
	}
}
