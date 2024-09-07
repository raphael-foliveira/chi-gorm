package database

import (
	"github.com/raphael-foliveira/chi-gorm/internal/config"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Start() error {
	db, err := start(config.DatabaseURL)
	if err != nil {
		return err
	}
	DB = db
	err = migrateDb()
	if err != nil {
		return err
	}
	return nil
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
	sqlDb, err := DB.DB()
	if err != nil {
		return err
	}
	if err := sqlDb.Close(); err != nil {
		return err
	}
	return nil
}
