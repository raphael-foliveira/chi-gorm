package service

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
)

type Products struct{}

func NewProducts() *Products {
	return &Products{}
}

func (c *Products) Create(schema *schemas.CreateProduct) (*entities.Product, error) {
	newProduct := schema.ToModel()
	err := productsRepository.Create(newProduct)
	return newProduct, err
}

func (c *Products) Update(id uint, schema *schemas.UpdateProduct) (*entities.Product, error) {
	entity, err := c.Get(id)
	if err != nil {
		return nil, err
	}
	c.updatePopulatedFields(entity, schema)
	err = productsRepository.Update(entity)
	return entity, err
}

func (c *Products) updatePopulatedFields(product *entities.Product, schema *schemas.UpdateProduct) {
	if schema.Name != "" {
		product.Name = schema.Name
	}
	if schema.Price != 0 {
		product.Price = schema.Price
	}
}

func (c *Products) Delete(id uint) error {
	client, err := c.Get(id)
	if err != nil {
		return err
	}
	err = productsRepository.Delete(client)
	if err != nil {
		return err
	}
	return nil
}

func (c *Products) List() ([]entities.Product, error) {
	return productsRepository.List()
}

func (c *Products) Get(id uint) (*entities.Product, error) {
	product, err := productsRepository.Get(id)
	if err != nil || product == nil {
		return nil, errProductNotFound
	}
	return product, nil
}
