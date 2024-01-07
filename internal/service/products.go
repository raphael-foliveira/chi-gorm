package service

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/exceptions"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
)

type Products interface {
	Create(schema *schemas.CreateProduct) (*entities.Product, error)
	Update(id uint, schema *schemas.UpdateProduct) (*entities.Product, error)
	Delete(id uint) error
	List() ([]entities.Product, error)
	Get(id uint) (*entities.Product, error)
}

type products struct {
	repository repository.Repository[entities.Product]
}

func NewProducts(repository repository.Repository[entities.Product]) Products {
	return &products{repository}
}

func (c *products) Create(schema *schemas.CreateProduct) (*entities.Product, error) {
	validationErr := schema.Validate()
	if validationErr != nil {
		return nil, validationErr
	}
	newProduct := schema.ToModel()
	err := c.repository.Create(newProduct)
	return newProduct, err
}

func (c *products) Update(id uint, schema *schemas.UpdateProduct) (*entities.Product, error) {
	validationErr := schema.Validate()
	if validationErr != nil {
		return nil, validationErr
	}
	entity, err := c.repository.Get(id)
	if err != nil {
		return nil, err
	}
	entity.Name = schema.Name
	entity.Price = schema.Price
	err = c.repository.Update(entity)
	return entity, err
}

func (c *products) Delete(id uint) error {
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

func (c *products) List() ([]entities.Product, error) {
	return c.repository.List()
}

func (c *products) Get(id uint) (*entities.Product, error) {
	product, err := c.repository.Get(id)
	if err != nil || product == nil {
		return nil, exceptions.NotFound("product not found")
	}
	return product, nil
}
