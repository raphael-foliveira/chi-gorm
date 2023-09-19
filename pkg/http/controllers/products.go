package controllers

import (
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/pkg/http/res"
	"github.com/raphael-foliveira/chi-gorm/pkg/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/pkg/repository"
)

type Products struct {
	productsRepo repository.Products
}

func NewProducts(r repository.Products) *Products {
	return &Products{r}
}

func (c *Products) Create(w http.ResponseWriter, r *http.Request) error {
	var body schemas.CreateProduct
	err := parseBody(r, &body)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest, "bad request")
	}
	newProduct := body.ToModel()
	err = c.productsRepo.Create(newProduct)
	if err != nil {
		return res.Error(w, err, http.StatusInternalServerError, "internal server error")
	}
	return res.JSON(w, http.StatusCreated, &newProduct)
}

func (c *Products) Update(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest, err.Error())
	}
	product, err := c.productsRepo.Get(id)
	if err != nil {
		return res.Error(w, err, http.StatusNotFound, "product not found")
	}
	var body schemas.UpdateProduct
	err = parseBody(r, &body)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest, err.Error())
	}
	product.Name = body.Name
	product.Price = body.Price
	err = c.productsRepo.Update(product)
	if err != nil {
		return res.Error(w, err, http.StatusInternalServerError, "internal server error")
	}
	return res.JSON(w, http.StatusOK, &product)
}

func (c *Products) Delete(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest, err.Error())
	}
	product, err := c.productsRepo.Get(id)
	if err != nil {
		return res.Error(w, err, http.StatusNotFound, "product not found")
	}
	err = c.productsRepo.Delete(product)
	if err != nil {
		return res.Error(w, err, http.StatusInternalServerError, "internal server error")
	}
	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (c *Products) List(w http.ResponseWriter, r *http.Request) error {
	products, err := c.productsRepo.List()
	if err != nil {
		return res.Error(w, err, http.StatusInternalServerError, "internal server error")
	}
	return res.JSON(w, http.StatusOK, &products)
}

func (c *Products) Get(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest, err.Error())
	}
	product, err := c.productsRepo.Get(id)
	if err != nil {
		return res.Error(w, err, http.StatusNotFound, "product not found")
	}
	return res.JSON(w, http.StatusOK, &product)
}
