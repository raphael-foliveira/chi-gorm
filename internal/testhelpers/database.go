package testhelpers

import (
	"log"

	"github.com/raphael-foliveira/chi-gorm/internal/database"
)

func StartDB() {
	if err := database.Start(); err != nil {
		log.Println("Error starting test database")
		log.Fatal(err)
	}
}
