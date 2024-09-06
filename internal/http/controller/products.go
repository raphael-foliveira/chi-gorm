package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

type Products struct {
	Service *service.Products
}

func NewProducts(service *service.Products) *Products {
	return &Products{service}
}

func (p *Products) Mount(mux *chi.Mux) {
	router := chi.NewRouter()
	router.Get("/", useHandler(p.List))
	router.Post("/", useHandler(p.Create))
	router.Get("/{id}", useHandler(p.Get))
	router.Delete("/{id}", useHandler(p.Delete))
	router.Put("/{id}", useHandler(p.Update))

	mux.Mount("/products", router)
}

func (p *Products) Create(ctx *Context) error {
	body := &schemas.CreateProduct{}
	err := ctx.ParseBody(body)
	if err != nil {
		return err
	}
	newOrder, err := p.Service.Create(body)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, schemas.NewProduct(newOrder))
}

func (p *Products) Update(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	body := &schemas.UpdateProduct{}
	err = ctx.ParseBody(body)
	if err != nil {
		return err
	}
	updatedOrder, err := p.Service.Update(id, body)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewProduct(updatedOrder))
}

func (p *Products) Delete(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	err = p.Service.Delete(id)
	if err != nil {
		return err
	}
	return ctx.SendStatus(http.StatusNoContent)
}

func (p *Products) List(ctx *Context) error {
	products, err := p.Service.List()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewProducts(products))
}

func (p *Products) Get(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	product, err := p.Service.Get(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewProduct(product))
}
