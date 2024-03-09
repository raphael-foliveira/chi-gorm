package repository

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/config"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
)

func TestMain(m *testing.M) {
	config.LoadCfg("../../.env.test")
	database.Db()
	m.Run()
	err := database.Close()
	if err != nil {
		panic(err)
	}
}
