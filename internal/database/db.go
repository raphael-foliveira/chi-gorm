package database

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initialize(databaseUrl string) *gorm.DB {
	db, err := start(databaseUrl)
	if err != nil {
		panic(err)
	}
	DB = db
	err = migrateDb()
	if err != nil {
		panic(err)
	}
	return DB
}

func start(dbUrl string) (*gorm.DB, error) {
	dialector := postgres.Open(dbUrl)
	return gorm.Open(dialector)
}

func migrateDb() error {
	return DB.AutoMigrate(
		&entities.Client{},
		&entities.Product{},
		&entities.Order{},
	)
}

func Close() error {
	if DB == nil {
		return nil
	}
	sqlDb, err := DB.DB()
	if err != nil {
		return err
	}
	err = sqlDb.Close()
	if err != nil {
		return err
	}
	DB = nil
	return nil
}
