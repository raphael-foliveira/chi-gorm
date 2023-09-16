package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/pkg/models"
	"github.com/raphael-foliveira/chi-gorm/pkg/repositories"
	"github.com/raphael-foliveira/chi-gorm/pkg/res"
	"github.com/raphael-foliveira/chi-gorm/pkg/schemas"
)

type Clients struct {
	repository repositories.Clients
}

func NewClients(r repositories.Clients) *Clients {
	return &Clients{r}
}

func (c *Clients) Create(w http.ResponseWriter, r *http.Request) error {
	body, err := c.parseCreate(w, r)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest, "bad request")
	}
	newClient := models.Client{
		Name:  body.Name,
		Email: body.Email,
	}
	err = c.repository.Create(&newClient)
	if err != nil {
		return res.Error(w, err, http.StatusInternalServerError, "internal server error")
	}
	return res.JSON(w, http.StatusCreated, &newClient)
}

func (c *Clients) Update(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest, "invalid user id")
	}
	client, err := c.repository.Get(id)
	if err != nil {
		return res.Error(w, err, http.StatusNotFound, "client not found")
	}
	body, err := c.parseUpdate(w, r)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest, "bad request")
	}
	client.Name = body.Name
	client.Email = body.Email
	err = c.repository.Update(client)
	if err != nil {
		return res.Error(w, err, http.StatusInternalServerError, "internal server error")
	}
	return res.JSON(w, http.StatusOK, &client)
}

func (c *Clients) Delete(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest, "invalid user id")
	}
	client, err := c.repository.Get(id)
	if err != nil {
		return res.Error(w, err, http.StatusNotFound, "client not found")
	}
	err = c.repository.Delete(client)
	if err != nil {
		return res.Error(w, err, http.StatusInternalServerError, "internal server error")
	}
	return res.SendStatus(w, http.StatusNoContent)
}

func (c *Clients) List(w http.ResponseWriter, r *http.Request) error {
	clients, err := c.repository.List()
	if err != nil {
		return res.Error(w, err, http.StatusInternalServerError, "internal server error")
	}
	return res.JSON(w, http.StatusOK, schemas.NewClients(clients))
}

func (c *Clients) Get(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest, "bad request")
	}
	client, err := c.repository.Get(id)
	if err != nil {
		return res.Error(w, err, http.StatusNotFound, "client not found")
	}
	return res.JSON(w, http.StatusOK, client)
}

func (c *Clients) parseCreate(w http.ResponseWriter, r *http.Request) (*schemas.CreateClient, error) {
	defer r.Body.Close()
	body := schemas.CreateClient{}
	return &body, json.NewDecoder(r.Body).Decode(&body)
}

func (c *Clients) parseUpdate(w http.ResponseWriter, r *http.Request) (*schemas.UpdateClient, error) {
	defer r.Body.Close()
	body := schemas.UpdateClient{}
	return &body, json.NewDecoder(r.Body).Decode(&body)
}
