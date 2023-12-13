package controllers

import (
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/internal/http/res"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/services"
)

type Products struct {
	service *services.Products
}

func NewProducts(service *services.Products) *Products {
	return &Products{service}
}

func (c *Products) Create(w http.ResponseWriter, r *http.Request) error {
	body, err := parseBody(r, &schemas.CreateProduct{})
	if err != nil {
		return err
	}
	newOrder, err := c.service.Create(body)
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusCreated, schemas.NewProduct(newOrder))
}

func (c *Products) Update(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return err
	}
	body, err := parseBody(r, &schemas.UpdateProduct{})
	if err != nil {
		return err
	}
	updatedOrder, err := c.service.Update(id, body)
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusOK, schemas.NewProduct(updatedOrder))
}

func (c *Products) Delete(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return err
	}
	err = c.service.Delete(id)
	if err != nil {
		return err
	}
	return res.SendStatus(w, http.StatusNoContent)
}

func (c *Products) List(w http.ResponseWriter, r *http.Request) error {
	products, err := c.service.List()
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusOK, schemas.NewProducts(products))
}

func (c *Products) Get(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return err
	}
	product, err := c.service.Get(id)
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusOK, schemas.NewProduct(product))
}
