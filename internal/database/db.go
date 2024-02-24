package database

import (
	"github.com/raphael-foliveira/chi-gorm/internal/cfg"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var instance *gorm.DB

func Db() *gorm.DB {
	if instance != nil {
		return instance
	}
	db, err := start(cfg.Cfg().DatabaseURL)
	if err != nil {
		panic(err)
	}
	instance = db
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
	return instance.AutoMigrate(&entities.Client{}, &entities.Product{}, &entities.Order{})
}

func CloseDb() error {
	if instance == nil {
		return nil
	}
	sqlDb, err := instance.DB()
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
