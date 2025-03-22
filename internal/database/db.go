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
	return migrateDb()
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
	return sqlDb.Close()
}
