package repository

import (
	"testing"
)

func TestProductsRepository(t *testing.T) {
	t.Run("Should find many", func(t *testing.T) {
		products, err := addProducts(2)
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
