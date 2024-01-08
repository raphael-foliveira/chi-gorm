package controller

import (
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/internal/http/res"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

type Clients struct {
	service *service.Clients
}

func NewClients(service *service.Clients) *Clients {
	return &Clients{service}
}

func (c *Clients) Create(w http.ResponseWriter, r *http.Request) error {
	body, err := parseBody(r, &schemas.CreateClient{})
	if err != nil {
		return err
	}
	newClient, err := c.service.Create(body)
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusCreated, &newClient)
}

func (c *Clients) Update(w http.ResponseWriter, r *http.Request) error {
	id, err := getUintPathParam(r, "id")
	if err != nil {
		return err
	}
	body, err := parseBody(r, &schemas.UpdateClient{})
	if err != nil {
		return err
	}
	updatedClient, err := c.service.Update(id, body)
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusOK, updatedClient)
}

func (c *Clients) Delete(w http.ResponseWriter, r *http.Request) error {
	id, err := getUintPathParam(r, "id")
	if err != nil {
		return err
	}
	err = c.service.Delete(id)
	if err != nil {
		return err
	}
	return res.SendStatus(w, http.StatusNoContent)
}

func (c *Clients) List(w http.ResponseWriter, r *http.Request) error {
	clients, err := c.service.List()
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusOK, schemas.NewClients(clients))
}

func (c *Clients) Get(w http.ResponseWriter, r *http.Request) error {
	id, err := getUintPathParam(r, "id")
	if err != nil {
		return err
	}
	client, err := c.service.Get(id)
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusOK, schemas.NewClientDetail(client))
}

func (c *Clients) GetProducts(w http.ResponseWriter, r *http.Request) error {
	id, err := getUintPathParam(r, "id")
	if err != nil {
		return err
	}
	orders, err := c.service.GetProducts(id)
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusOK, schemas.NewOrders(orders))
}
