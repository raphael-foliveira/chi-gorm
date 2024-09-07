package repository_test

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/config"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
)

func TestProductsRepository(t *testing.T) {
	config.Initialize("../../.env.test")
	database.Initialize(config.DatabaseURL)
	repository.Initialize()
	defer database.Close()

	t.Run("Should find many", func(t *testing.T) {
		products := []entities.Product{
			{
				Name:  "Brand 1",
				Price: 1.0,
			},
			{
				Name:  "Brand 2",
				Price: 2.0,
			},
		}
		database.DB.Create(&products)
		foundProducts, err := repository.Products.FindMany([]uint{products[0].ID, products[1].ID})
		if err != nil {
			t.Error(err)
		}
		if len(foundProducts) != 2 {
			t.Error("Should find 2 products")
		}
	})
}
