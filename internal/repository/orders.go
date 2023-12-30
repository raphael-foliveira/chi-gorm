package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
)

var Orders Repository[entities.Order] = &orders{&repository[entities.Order]{}}

type orders struct {
	*repository[entities.Order]
}
