package mocks

import (
	"errors"

	"github.com/raphael-foliveira/chi-gorm/internal/entities"
)

var ProductsStore = &productsStore{store[entities.Product]{}}

type productsStore struct {
	store[entities.Product]
}

func (cr *productsStore) FindMany(ids []uint) ([]entities.Product, error) {
	if cr.ShouldError {
		return nil, errors.New("")
	}
	products := []entities.Product{}
	for _, id := range ids {
		for _, product := range cr.Store {
			if product.ID == id {
				products = append(products, product)
			}
		}
	}
	return products, nil
}
