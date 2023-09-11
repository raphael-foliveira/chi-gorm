package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/raphael-foliveira/chi-gorm/pkg"
	"github.com/raphael-foliveira/chi-gorm/pkg/db"
	"github.com/raphael-foliveira/chi-gorm/pkg/server"
)

func main() {
	godotenv.Load()
	db := db.Connect(pkg.MainConfig.DatabaseURL)
	err := server.Start(db)
	if err != nil {
		fmt.Println(err)
	}
}
