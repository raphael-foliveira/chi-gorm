package main

import (
	"os"

	"github.com/raphael-foliveira/chi-gorm/internal/http/server"
	"gorm.io/driver/postgres"
)

func main() {
	dialector := postgres.Open(os.Getenv("DATABASE_URL"))
	err := server.NewServer(dialector).Start()
	if err != nil {
		panic(err)
	}

}
