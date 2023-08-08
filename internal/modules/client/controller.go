package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/pkg/res"
)

type Controller struct {
	repository iRepository
}

func NewController(r iRepository) *Controller {
	return &Controller{r}
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	newClient := Client{}
	err := json.NewDecoder(r.Body).Decode(&newClient)
	if err != nil {
		res.Error(w, 400, "bad request")
		return
	}
	defer r.Body.Close()
	err = c.repository.Create(&newClient)
	if err != nil {
		res.Error(w, 500, "internal server error")
		return
	}
	res.JSON(w, http.StatusCreated, &newClient)
}

func (c *Controller) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		res.Error(w, 400, "invalid user id")
		return
	}
	client, err := c.repository.Get(id)
	if err != nil {
		res.Error(w, 404, "client not found")
		return
	}
	err = json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		res.Error(w, 400, "bad request body")
		return
	}
	defer r.Body.Close()
	err = c.repository.Update(&client)
	if err != nil {
		res.Error(w, 500, "internal server error")
		return
	}
	res.JSON(w, http.StatusOK, &client)
}

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		res.Error(w, 400, "invalid user id")
		return
	}
	client, err := c.repository.Get(id)
	if err != nil {
		res.Error(w, 404, "client not found")
		return
	}
	err = c.repository.Delete(&client)
	if err != nil {
		res.Error(w, 500, "internal server error")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (c *Controller) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handling list")
	clients, err := c.repository.List()
	if err != nil {
		res.Error(w, 500, "internal server error")
		return
	}
	res.JSON(w, http.StatusOK, clients)
}

func (c *Controller) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		res.Error(w, 400, "bad request")
		return
	}
	client, err := c.repository.Get(id)
	if err != nil {
		res.Error(w, 404, "client not found")
		return
	}
	res.JSON(w, http.StatusOK, client)
}
