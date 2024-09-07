package service

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
)

type products struct{}

func NewProducts() *products {
	return &products{}
}

func (c *products) Create(schema *schemas.CreateProduct) (*entities.Product, error) {
	newProduct := schema.ToModel()
	err := repository.Products.Create(newProduct)
	return newProduct, err
}

func (c *products) Update(id uint, schema *schemas.UpdateProduct) (*entities.Product, error) {
	entity, err := c.Get(id)
	if err != nil {
		return nil, err
	}
	c.updatePopulatedFields(entity, schema)
	err = repository.Products.Update(entity)
	return entity, err
}

func (c *products) updatePopulatedFields(product *entities.Product, schema *schemas.UpdateProduct) {
	if schema.Name != "" {
		product.Name = schema.Name
	}
	if schema.Price != 0 {
		product.Price = schema.Price
	}
}

func (c *products) Delete(id uint) error {
	client, err := c.Get(id)
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
		return nil, errProductNotFound
	}
	return product, nil
}
