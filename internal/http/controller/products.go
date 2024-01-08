package controller

import (
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/internal/http/res"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

type Products struct {
	service *service.Products
}

func NewProducts(service *service.Products) *Products {
	return &Products{service}
}

func (p *Products) Create(w http.ResponseWriter, r *http.Request) error {
	body, err := parseBody(r, &schemas.CreateProduct{})
	if err != nil {
		return err
	}
	newOrder, err := p.service.Create(body)
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusCreated, schemas.NewProduct(newOrder))
}

func (p *Products) Update(w http.ResponseWriter, r *http.Request) error {
	id, err := getUintPathParam(r, "id")
	if err != nil {
		return err
	}
	body, err := parseBody(r, &schemas.UpdateProduct{})
	if err != nil {
		return err
	}
	updatedOrder, err := p.service.Update(id, body)
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusOK, schemas.NewProduct(updatedOrder))
}

func (p *Products) Delete(w http.ResponseWriter, r *http.Request) error {
	id, err := getUintPathParam(r, "id")
	if err != nil {
		return err
	}
	err = p.service.Delete(id)
	if err != nil {
		return err
	}
	return res.SendStatus(w, http.StatusNoContent)
}

func (p *Products) List(w http.ResponseWriter, r *http.Request) error {
	products, err := p.service.List()
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusOK, schemas.NewProducts(products))
}

func (p *Products) Get(w http.ResponseWriter, r *http.Request) error {
	id, err := getUintPathParam(r, "id")
	if err != nil {
		return err
	}
	product, err := p.service.Get(id)
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusOK, schemas.NewProduct(product))
}
