package order

import (
	"database/sql"
	"time"

	"github.com/raphael-foliveira/chi-gorm/internal/modules/client"
	"github.com/raphael-foliveira/chi-gorm/internal/modules/product"
)

type Order struct {
	ID        uint            `json:"id" gorm:"primarykey"`
	CreatedAt time.Time       `json:"-"`
	UpdatedAt time.Time       `json:"-"`
	DeletedAt sql.NullTime    `gorm:"index" json:"-"`
	ClientID  uint            `json:"-"`
	Client    client.Client   `json:"client"`
	ProductID uint            `json:"-"`
	Product   product.Product `json:"product"`
	Quantity  uint            `json:"quantity"`
}
