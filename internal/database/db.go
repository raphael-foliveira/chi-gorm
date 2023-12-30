package database

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb(dbUrl string) error {
	if Db != nil {
		return nil
	}
	dialector := postgres.Open(dbUrl)
	db, err := gorm.Open(dialector)
	if err != nil {
		return err
	}
	Db = db
	migrateDb()
	return nil
}

func migrateDb() error {
	return Db.AutoMigrate(&entities.Client{}, &entities.Product{}, &entities.Order{})
}

func CloseDb() error {
	sqlDb, err := Db.DB()
	if err != nil {
		return err
	}
	return sqlDb.Close()
}
