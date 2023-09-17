package db

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitPg() {
	if Db != nil {
		panic("Db already initialized")
	}
	godotenv.Load()
	var err error
	Db, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")))
	if err != nil {
		panic(err)
	}
}

func InitMemory() {
	if Db != nil {
		panic("Db already initialized")
	}
	var err error
	Db, err = gorm.Open(sqlite.Open(":memory:"))
	if err != nil {
		panic(err)
	}
}
