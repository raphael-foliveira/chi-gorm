package repository_test

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/config"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
)

func TestMain(m *testing.M) {
	config.Initialize("../../.env.test")
	database.Initialize(config.DatabaseURL)
	m.Run()
	database.Close()
}
