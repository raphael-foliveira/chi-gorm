package repository_test

import (
	"testing"

	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/testhelpers"
)

func TestProductsRepository(t *testing.T) {
	testhelpers.StartDB()
	defer database.Close()

	t.Run("Should find many", func(t *testing.T) {
		deps := newTestDependencies(t)
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
		foundProducts, err := deps.productsRepo.FindMany([]uint{products[0].ID, products[1].ID})
		if err != nil {
			t.Error(err)
		}
		if len(foundProducts) != 2 {
			t.Error("Should find 2 products")
		}
	})
}
