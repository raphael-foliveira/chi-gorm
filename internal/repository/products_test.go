package repository

import (
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/raphael-foliveira/chi-gorm/internal/cfg"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
)

func TestProductsRepository(t *testing.T) {
	cfg.LoadEnv("../../.env")
	database.InitDb(cfg.TestConfig.DatabaseURL)
	t.Run("Should find many", func(t *testing.T) {
		products := [2]entities.Product{}
		err := faker.FakeData(&products)
		if err != nil {
			t.Error(err)
		}
		for i := range products {
			products[i].ID = 0
		}
		err = database.Db.Create(&products).Error
		if err != nil {
			t.Error(err)
		}
		foundProducts, err := Products.FindMany([]uint{products[0].ID, products[1].ID})
		if err != nil {
			t.Error(err)
		}
		if len(foundProducts) != 2 {
			t.Error("Should find 2 products")
		}
	})

}