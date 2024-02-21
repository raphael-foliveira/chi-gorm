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
	config := cfg.GetCfg()
	db, err := database.GetDb(config.DatabaseURL)
	if err != nil {
		panic(err)
	}
	if err := server.NewApp(db).Start(); err != nil {
		panic(err)
	}
}
