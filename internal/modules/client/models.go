package client

import (
	"time"

	"github.com/raphael-foliveira/chi-gorm/internal/modules/order"
)

type Client struct {
	ID        uint          `json:"id" gorm:"primarykey" faker:"-"`
	CreatedAt time.Time     `json:"-" faker:"-"`
	UpdatedAt time.Time     `json:"-" faker:"-"`
	Name      string        `json:"name" gorm:"not null" faker:"name"`
	Email     string        `json:"email" gorm:"not null" faker:"email"`
	Orders    []order.Order `json:"orders" gorm:"OnDelete:CASCADE;"`
}
