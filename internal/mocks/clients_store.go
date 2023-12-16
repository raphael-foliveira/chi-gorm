package mocks

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
)

type ClientsStore struct {
	store[entities.Client]
}

func NewClientsStore() *ClientsStore {
	return &ClientsStore{newStore[entities.Client]()}
}
