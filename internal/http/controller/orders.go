package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
)

type Orders struct{}

func NewOrders() *Orders {
	return &Orders{}
}

func (o *Orders) Mount() {
	router := chi.NewRouter()
	router.Get("/", useHandler(o.List))
	router.Post("/", useHandler(o.Create))
	router.Get("/{id}", useHandler(o.Get))
	router.Delete("/{id}", useHandler(o.Delete))
	router.Put("/{id}", useHandler(o.Update))
	app.Mount("/orders", router)
}

func (o *Orders) Create(ctx *Context) error {
	body := &schemas.CreateOrder{}
	err := ctx.ParseBody(body)
	if err != nil {
		return err
	}
	newOrder, err := ordersService.Create(body)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, schemas.NewOrder(newOrder))
}

func (o *Orders) Update(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	body := &schemas.UpdateOrder{}
	err = ctx.ParseBody(body)
	if err != nil {
		return err
	}
	updatedOrder, err := ordersService.Update(id, body)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewOrder(updatedOrder))
}

func (o *Orders) Delete(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	err = ordersService.Delete(id)
	if err != nil {
		return err
	}
	return ctx.SendStatus(http.StatusNoContent)
}

func (o *Orders) List(ctx *Context) error {
	orders, err := ordersService.List()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewOrders(orders))
}

func (o *Orders) Get(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	order, err := ordersService.Get(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewOrder(order))
}
