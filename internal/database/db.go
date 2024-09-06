package database

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(databaseUrl string) *gorm.DB {
	db, err := start(databaseUrl)
	if err != nil {
		panic(err)
	}
	err = migrateDb(db)
	if err != nil {
		panic(err)
	}
	return db
}

func start(dbUrl string) (*gorm.DB, error) {
	dialector := postgres.Open(dbUrl)
	return gorm.Open(dialector)
}

func migrateDb(db *gorm.DB) error {
	return db.AutoMigrate(
		&entities.Client{},
		&entities.Product{},
		&entities.Order{},
	)
}

func Close(db *gorm.DB) {
	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDb.Close()
}
