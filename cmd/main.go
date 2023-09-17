package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/raphael-foliveira/chi-gorm/pkg/http/server"
	"github.com/raphael-foliveira/chi-gorm/pkg/persistence/store"
	"gorm.io/driver/postgres"
)

func main() {
	godotenv.Load()
	databaseUrl := os.Getenv("DATABASE_URL")
	gormDialector := postgres.Open(databaseUrl)
	store.InitSqlDb(gormDialector)
	err := server.Start()
	if err != nil {
		fmt.Println(err)
	}
}
