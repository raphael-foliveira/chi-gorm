package main

import (
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/http/server"
)

func main() {
	db := database.Db()
	defer database.Close()
	app := server.NewApp(db)
	if err := app.Start(); err != nil {
		panic(err)
	}

}
