package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
)

type Clients struct{}

func NewClients() *Clients {
	return &Clients{}
}

func (c *Clients) Mount() {
	router := chi.NewRouter()
	router.Get("/", useHandler(c.List))
	router.Get("/{id}", useHandler(c.Get))
	router.Get("/{id}/products", useHandler(c.GetProducts))
	router.Post("/", useHandler(c.Create))
	router.Delete("/{id}", useHandler(c.Delete))
	router.Put("/{id}", useHandler(c.Update))

	app.Mount("/clients", router)
}

func (c *Clients) Create(ctx *Context) error {
	body := &schemas.CreateClient{}
	err := ctx.ParseBody(body)
	if err != nil {
		return err
	}
	newClient, err := clientsService.Create(body)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, &newClient)
}

func (c *Clients) Update(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	body := &schemas.UpdateClient{}
	err = ctx.ParseBody(body)
	if err != nil {
		return err
	}
	updatedClient, err := clientsService.Update(id, body)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, updatedClient)
}

func (c *Clients) Delete(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	err = clientsService.Delete(id)
	if err != nil {
		return err
	}
	return ctx.SendStatus(http.StatusNoContent)
}

func (c *Clients) List(ctx *Context) error {
	clients, err := clientsService.List()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewClients(clients))
}

func (c *Clients) Get(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	client, err := clientsService.Get(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewClientDetail(client))
}

func (c *Clients) GetProducts(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	orders, err := clientsService.GetOrders(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewOrders(orders))
}
