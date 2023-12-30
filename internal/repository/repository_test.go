package repository

import (
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/raphael-foliveira/chi-gorm/internal/cfg"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
)

func TestMain(m *testing.M) {
	err := cfg.LoadCfg("../../.env")
	if err != nil {
		panic(err)
	}
	err = database.InitDb(cfg.TestConfig.DatabaseURL)
	if err != nil {
		panic(err)
	}
	m.Run()
	err = database.CloseDb()
	if err != nil {
		panic(err)
	}
}

func addProducts(quantity int) ([]entities.Product, error) {
	products := []entities.Product{}
	err := faker.FakeData(&products)
	if err != nil {
		return nil, err
	}
	for i := range products {
		products[i].ID = 0
	}
	return products, database.Db.Create(&products).Error
}
