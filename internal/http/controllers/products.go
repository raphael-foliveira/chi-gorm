package controllers

import (
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/internal/http/res"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
)

type Products struct {
	repository repository.Products
}

func NewProducts(productsRepo repository.Products) *Products {
	return &Products{productsRepo}
}

func (c *Products) Create(w http.ResponseWriter, r *http.Request) error {
	var body schemas.CreateProduct
	err := parseBody(r, &body)
	if err != nil {
		return err
	}
	newProduct := body.ToModel()
	err = c.repository.Create(newProduct)
	if err != nil {
		return res.Error(w, http.StatusInternalServerError, err.Error())
	}
	return res.JSON(w, http.StatusCreated, schemas.NewProduct(newProduct))
}

func (c *Products) Update(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return err
	}
	product, err := c.repository.Get(id)
	if err != nil {
		return res.Error(w, http.StatusNotFound, err.Error())
	}
	var body schemas.UpdateProduct
	err = parseBody(r, &body)
	if err != nil {
		return err
	}
	product.Name = body.Name
	product.Price = body.Price
	err = c.repository.Update(product)
	if err != nil {
		return res.Error(w, http.StatusInternalServerError, err.Error())
	}
	return res.JSON(w, http.StatusOK, schemas.NewProduct(product))
}

func (c *Products) Delete(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return res.Error(w, http.StatusInternalServerError, err.Error())
	}
	product, err := c.repository.Get(id)
	if err != nil {
		return res.Error(w, http.StatusNotFound, err.Error())
	}
	err = c.repository.Delete(product)
	if err != nil {
		return res.Error(w, http.StatusInternalServerError, err.Error())
	}
	return res.SendStatus(w, http.StatusNoContent)
}

func (c *Products) List(w http.ResponseWriter, r *http.Request) error {
	products, err := c.repository.List()
	if err != nil {
		return res.Error(w, http.StatusInternalServerError, err.Error())
	}
	return res.JSON(w, http.StatusOK, schemas.NewProducts(products))
}

func (c *Products) Get(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return err
	}
	product, err := c.repository.Get(id)
	if err != nil {
		return res.Error(w, http.StatusNotFound, err.Error())
	}
	return res.JSON(w, http.StatusOK, schemas.NewProduct(product))
}
