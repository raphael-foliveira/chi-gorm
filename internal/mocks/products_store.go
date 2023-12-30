package mocks

import (
	"errors"

	"github.com/raphael-foliveira/chi-gorm/internal/entities"
)

var ProductsStore = &ProductsStoreImpl{store[entities.Product]{}}

type ProductsStoreImpl struct {
	store[entities.Product]
}

func (cr *ProductsStoreImpl) FindMany(ids []uint) ([]entities.Product, error) {
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
