package controller

import (
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/internal/http/res"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

var Clients = &clients{}

type clients struct{}

func (c *clients) Create(w http.ResponseWriter, r *http.Request) error {
	body, err := parseBody(r, &schemas.CreateClient{})
	if err != nil {
		return err
	}
	newClient, err := service.Clients.Create(body)
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusCreated, &newClient)
}

func (c *clients) Update(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return err
	}
	body, err := parseBody(r, &schemas.UpdateClient{})
	if err != nil {
		return err
	}
	updatedClient, err := service.Clients.Update(id, body)
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusOK, updatedClient)
}

func (c *clients) Delete(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return err
	}
	err = service.Clients.Delete(id)
	if err != nil {
		return err
	}
	return res.SendStatus(w, http.StatusNoContent)
}

func (c *clients) List(w http.ResponseWriter, r *http.Request) error {
	clients, err := service.Clients.List()
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusOK, schemas.NewClients(clients))
}

func (c *clients) Get(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return err
	}
	client, err := service.Clients.Get(id)
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusOK, schemas.NewClientDetail(client))
}
