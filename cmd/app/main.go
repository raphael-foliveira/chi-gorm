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
	db, err := database.GetDb(cfg.Cfg.DatabaseURL)
	if err != nil {
		panic(err)
	}
	if err := server.NewServer(db).Start(); err != nil {
		panic(err)
	}
}
