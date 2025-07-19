package api

import (
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/ports"
)

type ClientsController struct {
	clientsRepo ports.ClientsRepository
	ordersRepo  ports.OrdersRepository
}

func NewClientsController(clientsRepo ports.ClientsRepository, ordersRepo ports.OrdersRepository) *ClientsController {
	return &ClientsController{
		clientsRepo: clientsRepo,
		ordersRepo:  ordersRepo,
	}
}

func (c *ClientsController) Create(ctx *Context) error {
	var body schemas.CreateClient
	err := ctx.ParseBody(&body)
	if err != nil {
		return err
	}
	newClient := body.ToModel()
	if err := c.clientsRepo.Create(newClient); err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, &newClient)
}

func (c *ClientsController) Update(ctx *Context) error {
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

func (c *ClientsController) Delete(ctx *Context) error {
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

func (c *ClientsController) List(ctx *Context) error {
	clients, err := c.clientsRepo.List()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewClients(clients))
}

func (c *ClientsController) Get(ctx *Context) error {
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

func (c *ClientsController) GetProducts(ctx *Context) error {
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
