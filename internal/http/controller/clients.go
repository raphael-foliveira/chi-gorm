package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/ports"
)

type Clients struct {
	clientsRepo ports.ClientsRepository
	ordersRepo  ports.OrdersRepository
}

func NewClients(clientsRepo ports.ClientsRepository, ordersRepo ports.OrdersRepository) *Clients {
	return &Clients{
		clientsRepo: clientsRepo,
		ordersRepo:  ordersRepo,
	}
}

func (c *Clients) Mount(mux *chi.Mux) {
	router := NewRouter()
	router.Get("/", c.List)
	router.Get("/{id}", c.Get)
	router.Get("/{id}/products", c.GetProducts)
	router.Post("/", c.Create)
	router.Delete("/{id}", c.Delete)
	router.Put("/{id}", c.Update)
	mux.Mount("/clients", router)
}

func (c *Clients) Create(ctx *Context) error {
	var body schemas.CreateClient
	err := ctx.ParseBody(&body)
	if err != nil {
		return err
	}
	newClient := body.ToModel()
	if err := c.clientsRepo.Create(body.ToModel()); err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, &newClient)
}

func (c *Clients) Update(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	var body schemas.UpdateClient
	err = ctx.ParseBody(&body)
	if err != nil {
		return err
	}

	client := body.ToModel()
	client.ID = id

	if err := c.clientsRepo.Update(client); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, client)
}

func (c *Clients) Delete(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	err = c.clientsRepo.Delete(id)
	if err != nil {
		return err
	}
	return ctx.SendStatus(http.StatusNoContent)
}

func (c *Clients) List(ctx *Context) error {
	clients, err := c.clientsRepo.List()
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
	client, err := c.clientsRepo.Get(id)
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
	orders, err := c.ordersRepo.FindByClient(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewOrders(orders))
}
