package controllers

import (
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/pkg/http/res"
	"github.com/raphael-foliveira/chi-gorm/pkg/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/pkg/models"
	"github.com/raphael-foliveira/chi-gorm/pkg/persistence/sqlstore"
)

type Products struct {
	productsStore sqlstore.Products
}

func NewProducts(r sqlstore.Products) *Products {
	return &Products{r}
}

func (c *Products) Create(w http.ResponseWriter, r *http.Request) error {
	var body schemas.CreateProduct
	err := parseBody(r, &body)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest, "bad request")
	}
	newProduct := models.Product{
		Name:  body.Name,
		Price: body.Price,
	}
	err = c.productsStore.Create(&newProduct)
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
	product, err := c.productsStore.Get(id)
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
	err = c.productsStore.Update(product)
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
	product, err := c.productsStore.Get(id)
	if err != nil {
		return res.Error(w, err, http.StatusNotFound, "product not found")
	}
	err = c.productsStore.Delete(product)
	if err != nil {
		return res.Error(w, err, http.StatusInternalServerError, "internal server error")
	}
	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (c *Products) List(w http.ResponseWriter, r *http.Request) error {
	products, err := c.productsStore.List()
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
	product, err := c.productsStore.Get(id)
	if err != nil {
		return res.Error(w, err, http.StatusNotFound, "product not found")
	}
	return res.JSON(w, http.StatusOK, &product)
}
