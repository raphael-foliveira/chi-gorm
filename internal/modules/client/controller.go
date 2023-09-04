package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/interfaces"
	"github.com/raphael-foliveira/chi-gorm/pkg/res"
)

type Controller struct {
	repository interfaces.IRepository[Client]
}

func NewController(r interfaces.IRepository[Client]) *Controller {
	return &Controller{r}
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) error {
	newClient := Client{}
	err := json.NewDecoder(r.Body).Decode(&newClient)
	if err != nil {
		return res.Error(w, 400, "bad request", err)
	}
	defer r.Body.Close()
	err = c.repository.Create(&newClient)
	if err != nil {
		return res.Error(w, 500, "internal server error", err)
	}
	return res.JSON(w, http.StatusCreated, &newClient)
}

func (c *Controller) Update(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return res.Error(w, 400, "invalid user id", err)
	}
	client, err := c.repository.Get(id)
	if err != nil {
		return res.Error(w, 404, "client not found", err)
	}
	err = json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		return res.Error(w, 400, "bad request body", err)
	}
	defer r.Body.Close()
	err = c.repository.Update(&client)
	if err != nil {
		return res.Error(w, 500, "internal server error", err)
	}
	return res.JSON(w, http.StatusOK, &client)
}

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return res.Error(w, 400, "invalid user id", err)
	}
	client, err := c.repository.Get(id)
	if err != nil {
		return res.Error(w, 404, "client not found", err)
	}
	err = c.repository.Delete(&client)
	if err != nil {
		return res.Error(w, 500, "internal server error", err)
	}
	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (c *Controller) List(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("handling list")
	clients, err := c.repository.List()
	if err != nil {
		return res.Error(w, 500, "internal server error", err)
	}
	return res.JSON(w, http.StatusOK, clients)
}

func (c *Controller) Get(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return res.Error(w, 400, "bad request", err)
	}
	client, err := c.repository.Get(id)
	if err != nil {
		return res.Error(w, 404, "client not found", err)
	}
	return res.JSON(w, http.StatusOK, client)
}
