package service

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/config"
)

func TestMain(m *testing.M) {
	config.LoadCfg("../../.env.test")
	m.Run()
}
