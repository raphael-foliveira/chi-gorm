package mocks

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
)

var ClientsStore = &clientsStore{store[entities.Client]{}}

type clientsStore struct {
	store[entities.Client]
}
