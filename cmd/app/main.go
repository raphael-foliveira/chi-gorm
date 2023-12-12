package main

import (
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/http/server"
)

func main() {
	database.InitPg()
	err := server.Start()
	if err != nil {
		panic(err)
	}
}
