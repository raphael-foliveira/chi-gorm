package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/raphael-foliveira/chi-gorm/internal/db"
	"github.com/raphael-foliveira/chi-gorm/internal/server"
)

func main() {
	godotenv.Load()
	db := db.Connect(os.Getenv("DATABASE_URL"))
	server.Start(db)
}
