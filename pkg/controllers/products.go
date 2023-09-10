package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/pkg/interfaces"
	"github.com/raphael-foliveira/chi-gorm/pkg/models"
	"github.com/raphael-foliveira/chi-gorm/pkg/res"
	"github.com/raphael-foliveira/chi-gorm/pkg/schemas"
)

type Products struct {
	repository interfaces.Repository[models.Product]
}

func NewProducts(r interfaces.Repository[models.Product]) *Products {
	return &Products{r}
}

func (c *Products) Create(w http.ResponseWriter, r *http.Request) error {
	body, err := c.parseCreate(w, r)
	if err != nil {
		return res.New(w).Status(http.StatusBadRequest).Error("bad request")
	}
	newProduct := models.Product{
		Name:  body.Name,
		Price: body.Price,
	}
	err = c.repository.Create(&newProduct)
	if err != nil {
		return res.New(w).Status(http.StatusInternalServerError).Error("internal server error")
	}
	return res.New(w).Status(http.StatusCreated).JSON(&newProduct)
}

func (c *Products) Update(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return res.New(w).Status(http.StatusBadRequest).Error("bad request")
	}
	product, err := c.repository.Get(id)
	if err != nil {
		return res.New(w).Status(http.StatusNotFound).Error("product not found")
	}
	body, err := c.parseUpdate(w, r)
	if err != nil {
		return res.New(w).Status(http.StatusBadRequest).Error("bad request")
	}
	product.Name = body.Name
	product.Price = body.Price
	err = c.repository.Update(&product)
	if err != nil {
		return res.New(w).Status(http.StatusInternalServerError).Error("internal server error")
	}
	return res.New(w).JSON(&product)
}

func (c *Products) Delete(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return res.New(w).Status(http.StatusBadRequest).Error("bad request")
	}
	product, err := c.repository.Get(id)
	if err != nil {
		return res.New(w).Status(http.StatusNotFound).Error("product not found")
	}
	err = c.repository.Delete(&product)
	if err != nil {
		return res.New(w).Status(http.StatusInternalServerError).Error("internal server error")
	}
	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (c *Products) List(w http.ResponseWriter, r *http.Request) error {
	products, err := c.repository.List()
	if err != nil {
		return res.New(w).Status(http.StatusInternalServerError).Error("internal server error")
	}
	return res.New(w).JSON(&products)
}

func (c *Products) Get(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return res.New(w).Status(http.StatusBadRequest).Error("bad request")
	}
	product, err := c.repository.Get(id)
	if err != nil {
		return res.New(w).Status(http.StatusNotFound).Error("product not found")
	}
	return res.New(w).JSON(&product)
}

func (c *Products) parseCreate(w http.ResponseWriter, r *http.Request) (*schemas.CreateProduct, error) {
	defer r.Body.Close()
	body := schemas.CreateProduct{}
	return &body, json.NewDecoder(r.Body).Decode(&body)
}

func (c *Products) parseUpdate(w http.ResponseWriter, r *http.Request) (*schemas.UpdateProduct, error) {
	defer r.Body.Close()
	body := schemas.UpdateProduct{}
	return &body, json.NewDecoder(r.Body).Decode(&body)
}
