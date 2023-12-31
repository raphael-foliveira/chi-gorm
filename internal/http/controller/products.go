package controller

import (
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/internal/http/res"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

var Products = &products{}

type products struct{}

func (c *products) Create(w http.ResponseWriter, r *http.Request) error {
	body, err := parseBody(r, &schemas.CreateProduct{})
	if err != nil {
		return err
	}
	newOrder, err := service.Products.Create(body)
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusCreated, schemas.NewProduct(newOrder))
}

func (c *products) Update(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return err
	}
	body, err := parseBody(r, &schemas.UpdateProduct{})
	if err != nil {
		return err
	}
	updatedOrder, err := service.Products.Update(id, body)
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusOK, schemas.NewProduct(updatedOrder))
}

func (c *products) Delete(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return err
	}
	err = service.Products.Delete(id)
	if err != nil {
		return err
	}
	return res.SendStatus(w, http.StatusNoContent)
}

func (c *products) List(w http.ResponseWriter, r *http.Request) error {
	products, err := service.Products.List()
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusOK, schemas.NewProducts(products))
}

func (c *products) Get(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return err
	}
	product, err := service.Products.Get(id)
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusOK, schemas.NewProduct(product))
}
