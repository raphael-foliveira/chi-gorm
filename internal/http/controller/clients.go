package controller

import (
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

type clients struct{}

func NewClients() *clients {
	return &clients{}
}

func (c *clients) Create(ctx *Context) error {
	body := &schemas.CreateClient{}
	err := ctx.ParseBody(body)
	if err != nil {
		return err
	}
	newClient, err := service.Clients.Create(body)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, &newClient)
}

func (c *clients) Update(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	body := &schemas.UpdateClient{}
	err = ctx.ParseBody(body)
	if err != nil {
		return err
	}
	updatedClient, err := service.Clients.Update(id, body)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, updatedClient)
}

func (c *clients) Delete(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	err = service.Clients.Delete(id)
	if err != nil {
		return err
	}
	return ctx.SendStatus(http.StatusNoContent)
}

func (c *clients) List(ctx *Context) error {
	clients, err := service.Clients.List()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewClients(clients))
}

func (c *clients) Get(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	client, err := service.Clients.Get(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewClientDetail(client))
}

func (c *clients) GetProducts(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	orders, err := service.Clients.GetOrders(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewOrders(orders))
}
