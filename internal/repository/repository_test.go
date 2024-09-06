package repository

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/config"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
)

func TestMain(m *testing.M) {
	config := config.LoadCfg("../../.env.test")
	db := database.New(config.DatabaseURL)
	m.Run()
	database.Close(db)
}
