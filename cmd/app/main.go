package main

import (
	"github.com/raphael-foliveira/chi-gorm/internal/config"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/http/server"
)

func main() {
	config := config.Config()
	db := database.Db(config.DatabaseURL)
	defer database.Close()
	app := server.NewApp(db)
	if err := app.Start(); err != nil {
		panic(err)
	}
}
