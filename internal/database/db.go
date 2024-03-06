package database

import (
	"github.com/raphael-foliveira/chi-gorm/internal/config"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

var instance *DB

func Db() *DB {
	if instance != nil {
		return instance
	}
	db, err := start(config.Cfg().DatabaseURL)
	if err != nil {
		panic(err)
	}
	instance = &DB{db}
	err = migrateDb()
	if err != nil {
		panic(err)
	}
	return instance
}

func start(dbUrl string) (*gorm.DB, error) {
	dialector := postgres.Open(dbUrl)
	return gorm.Open(dialector)
}

func migrateDb() error {
	return instance.AutoMigrate(
		&entities.Client{},
		&entities.Product{},
		&entities.Order{},
	)
}

func CloseDb() error {
	if instance == nil {
		return nil
	}
	sqlDb, err := instance.DB.DB()
	if err != nil {
		return err
	}
	err = sqlDb.Close()
	if err != nil {
		return err
	}
	instance = nil
	return nil
}
