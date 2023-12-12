package controllers

import (
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/internal/http/res"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
)

type Clients struct {
	repository repository.Clients
}

func NewClients(clientsRepo repository.Clients) *Clients {
	return &Clients{clientsRepo}
}

func (c *Clients) Create(w http.ResponseWriter, r *http.Request) error {
	var body schemas.CreateClient
	err := parseBody(r, &body)
	if err != nil {
		return err
	}
	newClient := body.ToModel()
	err = c.repository.Create(&newClient)
	if err != nil {
		return res.Error(w, http.StatusInternalServerError, err.Error())
	}
	return res.JSON(w, http.StatusCreated, &newClient)
}

func (c *Clients) Update(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return err
	}
	client, err := c.repository.Get(id)
	if err != nil {
		return res.Error(w, http.StatusNotFound, err.Error())
	}
	var body schemas.UpdateClient
	err = parseBody(r, &body)
	if err != nil {
		return err
	}
	client.Name = body.Name
	client.Email = body.Email
	err = c.repository.Update(client)
	if err != nil {
		return res.Error(w, http.StatusInternalServerError, err.Error())
	}
	return res.JSON(w, http.StatusOK, &client)
}

func (c *Clients) Delete(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return err
	}
	client, err := c.repository.Get(id)
	if err != nil {
		return res.Error(w, http.StatusNotFound, err.Error())
	}
	err = c.repository.Delete(client)
	if err != nil {
		return res.Error(w, http.StatusInternalServerError, err.Error())
	}
	return res.SendStatus(w, http.StatusNoContent)
}

func (c *Clients) List(w http.ResponseWriter, r *http.Request) error {
	clients, err := c.repository.List()
	if err != nil {
		return res.Error(w, http.StatusInternalServerError, err.Error())
	}
	return res.JSON(w, http.StatusOK, schemas.NewClients(clients))
}

func (c *Clients) Get(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return err
	}
	client, err := c.repository.Get(id)
	if err != nil {
		return res.Error(w, http.StatusNotFound, err.Error())
	}
	return res.JSON(w, http.StatusOK, schemas.NewClientDetail(client))
}
