package main

import (
	"github.com/raphael-foliveira/chi-gorm/internal/cfg"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/http/server"
)

func main() {
	if err := cfg.LoadCfg(".env"); err != nil {
		panic(err)
	}
	if err := database.InitDb(cfg.DatabaseURL); err != nil {
		panic(err)
	}
	if err := server.NewServer().Start(); err != nil {
		panic(err)
	}
}
