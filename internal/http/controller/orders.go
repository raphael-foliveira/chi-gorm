package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/ports"
)

type Orders struct {
	ordersRepo ports.OrdersRepository
}

func NewOrders(ordersRepo ports.OrdersRepository) *Orders {
	return &Orders{
		ordersRepo: ordersRepo,
	}
}

func (c *Orders) Mount(mux chi.Router) {
	router := NewRouter()
	router.Get("/", c.List)
	router.Post("/", c.Create)
	router.Get("/{id}", c.Get)
	router.Delete("/{id}", c.Delete)
	router.Put("/{id}", c.Update)
	mux.Mount("/orders", router)
}

func (o *Orders) Create(ctx *Context) error {
	var body schemas.CreateOrder
	err := ctx.ParseBody(&body)
	if err != nil {
		return err
	}
	newOrder := body.ToModel()
	if err := o.ordersRepo.Create(newOrder); err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, schemas.NewOrder(newOrder))
}

func (o *Orders) Update(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	var body schemas.UpdateOrder
	err = ctx.ParseBody(&body)
	if err != nil {
		return err
	}
	order := body.ToModel()
	order.ID = id
	if err := o.ordersRepo.Update(order); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewOrder(order))
}

func (o *Orders) Delete(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	err = o.ordersRepo.Delete(id)
	if err != nil {
		return err
	}
	return ctx.SendStatus(http.StatusNoContent)
}

func (o *Orders) List(ctx *Context) error {
	orders, err := o.ordersRepo.List()
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
	order, err := o.ordersRepo.Get(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewOrder(order))
}
