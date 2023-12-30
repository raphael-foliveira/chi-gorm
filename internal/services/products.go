package services

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/exceptions"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
)

var productNotFoundErr = &exceptions.NotFoundError{Entity: "product"}
var Products = &products{}

type products struct{}

func (c *products) Create(schema *schemas.CreateProduct) (*entities.Product, error) {
	validationErr := schema.Validate()
	if validationErr != nil {
		return nil, validationErr
	}
	newProduct := schema.ToModel()
	err := repository.Products.Create(newProduct)
	return newProduct, err
}

func (c *products) Update(id uint, schema *schemas.UpdateProduct) (*entities.Product, error) {
	validationErr := schema.Validate()
	if validationErr != nil {
		return nil, validationErr
	}
	entity, err := repository.Products.Get(id)
	if err != nil {
		return nil, err
	}
	entity.Name = schema.Name
	entity.Price = schema.Price
	err = repository.Products.Update(entity)
	return entity, err
}

func (c *products) Delete(id uint) error {
	client, err := repository.Products.Get(id)
	if err != nil {
		return err
	}
	err = repository.Products.Delete(client)
	if err != nil {
		return err
	}
	return nil
}

func (c *products) List() ([]entities.Product, error) {
	return repository.Products.List()
}

func (c *products) Get(id uint) (*entities.Product, error) {
	product, err := repository.Products.Get(id)
	if err != nil || product == nil {
		return nil, productNotFoundErr
	}
	return product, nil
}
