package service_test

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/config"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

func TestMain(m *testing.M) {
	config.Initialize("../../.env.test")
	service.Initialize()
	m.Run()
}
