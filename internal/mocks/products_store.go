package mocks

import (
	"errors"

	"github.com/raphael-foliveira/chi-gorm/internal/entities"
)

type ProductsStore struct {
	store[entities.Product]
}

func NewProductsStore() *ProductsStore {
	return &ProductsStore{newStore[entities.Product]()}
}

func (cr *ProductsStore) FindMany(ids []int64) ([]entities.Product, error) {
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
