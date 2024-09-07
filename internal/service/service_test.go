package service_test

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/config"
)

func TestMain(m *testing.M) {
	config.Load("../../.env.test")
	m.Run()
}
