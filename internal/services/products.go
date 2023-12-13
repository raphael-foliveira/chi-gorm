package services

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/exceptions"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
)

var productNotFoundErr = &exceptions.NotFoundError{Entity: "product"}

type Products struct {
	repository repository.Products
}

func NewProducts(productsRepo repository.Products) *Products {
	return &Products{productsRepo}
}

func (c *Products) Create(schema *schemas.CreateProduct) (*entities.Product, error) {
	newProduct := schema.ToModel()
	err := c.repository.Create(newProduct)
	return newProduct, err
}

func (c *Products) Update(id int64, schema *schemas.UpdateProduct) (*entities.Product, error) {
	entity, err := c.repository.Get(id)
	if err != nil {
		return nil, err
	}
	entity.Name = schema.Name
	entity.Price = schema.Price
	err = c.repository.Update(entity)
	return entity, err
}

func (c *Products) Delete(id int64) error {
	client, err := c.repository.Get(id)
	if err != nil {
		return err
	}
	err = c.repository.Delete(client)
	if err != nil {
		return err
	}
	return nil
}

func (c *Products) List() ([]entities.Product, error) {
	return c.repository.List()
}

func (c *Products) Get(id int64) (*entities.Product, error) {
	client, err := c.repository.Get(id)
	if err != nil {
		return nil, productNotFoundErr
	}
	return client, nil
}
