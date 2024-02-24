package main

import (
	"github.com/raphael-foliveira/chi-gorm/internal/http/server"
)

func main() {
	app := server.App()
	if err := app.Start(); err != nil {
		panic(err)
	}
}
