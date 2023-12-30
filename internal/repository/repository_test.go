package repository

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/cfg"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
)

func TestMain(m *testing.M) {
	cfg.LoadCfg("../../.env")
	database.InitDb(cfg.TestConfig.DatabaseURL)
	m.Run()
	database.CloseDb()
}
