package service

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
)

type ProductsService struct {
	repository repository.ProductsRepository
}

func NewProductsService(repository repository.ProductsRepository) *ProductsService {
	return &ProductsService{repository}
}

func (c *ProductsService) Create(schema *schemas.CreateProduct) (*entities.Product, error) {
	newProduct := schema.ToModel()
	err := c.repository.Create(newProduct)
	return newProduct, err
}

func (c *ProductsService) Update(id uint, schema *schemas.UpdateProduct) (*entities.Product, error) {
	entity, err := c.Get(id)
	if err != nil {
		return nil, err
	}
	entity.Name = schema.Name
	entity.Price = schema.Price
	err = c.repository.Update(entity)
	return entity, err
}

func (c *ProductsService) Delete(id uint) error {
	client, err := c.Get(id)
	if err != nil {
		return err
	}
	err = c.repository.Delete(client)
	if err != nil {
		return err
	}
	return nil
}

func (c *ProductsService) List() ([]entities.Product, error) {
	return c.repository.List()
}

func (c *ProductsService) Get(id uint) (*entities.Product, error) {
	product, err := c.repository.Get(id)
	if err != nil || product == nil {
		return nil, errProductNotFound
	}
	return product, nil
}
