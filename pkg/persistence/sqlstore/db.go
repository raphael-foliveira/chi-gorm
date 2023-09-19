package sqlstore

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitPg() {
	if db != nil {
		panic("db already initialized")
	}
	godotenv.Load()
	var err error
	db, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")))
	if err != nil {
		panic(err)
	}
}

func InitMemory() {
	if db != nil {
		panic("db already initialized")
	}
	var err error
	db, err = gorm.Open(sqlite.Open(":memory:"))
	if err != nil {
		panic(err)
	}
}

func GetInstance() *gorm.DB {
	return db
}

func CloseDb() {
	sqlDb, _ := db.DB()
	sqlDb.Close()
	db = nil
}
