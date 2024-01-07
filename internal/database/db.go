package database

import (
	"fmt"

	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var instance *gorm.DB

func GetDb(dbUrl string) (db *gorm.DB, err error) {
	if instance != nil {
		return instance, nil
	}
	dialector := postgres.Open(dbUrl)
	db, err = gorm.Open(dialector)
	if err != nil {
		return nil, err
	}
	instance = db
	migrateDb()
	return instance, nil
}

func GetInstance() *gorm.DB {
	return instance
}

func migrateDb() error {
	return instance.AutoMigrate(&entities.Client{}, &entities.Product{}, &entities.Order{})
}

func CloseDb() error {
	sqlDb, err := instance.DB()
	if err != nil {
		return err
	}
	err = sqlDb.Close()
	if err != nil {
		return err
	}
	fmt.Println("instance:", instance)
	instance = nil
	fmt.Println("instance:", instance)
	return nil
}
