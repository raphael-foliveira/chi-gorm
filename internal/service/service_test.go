package service

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/cfg"
)

func TestMain(m *testing.M) {
	cfg.LoadCfg("../../.env.test")
	m.Run()
}
