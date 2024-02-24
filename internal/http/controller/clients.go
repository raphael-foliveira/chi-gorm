package controller

import (
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/internal/http/res"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

type clients struct {
	service *service.ClientsService
}

func NewClients(service *service.ClientsService) *clients {
	return &clients{service}
}

func (c *clients) Create(w http.ResponseWriter, r *http.Request) error {
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

func (c *clients) Update(w http.ResponseWriter, r *http.Request) error {
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

func (c *clients) Delete(w http.ResponseWriter, r *http.Request) error {
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

func (c *clients) List(w http.ResponseWriter, r *http.Request) error {
	clients, err := c.service.List()
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusOK, schemas.NewClients(clients))
}

func (c *clients) Get(w http.ResponseWriter, r *http.Request) error {
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

func (c *clients) GetProducts(w http.ResponseWriter, r *http.Request) error {
	id, err := getUintPathParam(r, "id")
	if err != nil {
		return err
	}
	orders, err := c.service.GetOrders(id)
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusOK, schemas.NewOrders(orders))
}
