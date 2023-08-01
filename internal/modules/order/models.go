package order

import (
	"database/sql"
	"time"

	"github.com/raphael-foliveira/chi-gorm/internal/modules/client"
	"github.com/raphael-foliveira/chi-gorm/internal/modules/product"
)

type Order struct {
	ID        uint            `json:"id" gorm:"primarykey" faker:"-"`
	CreatedAt time.Time       `json:"-" faker:"-"`
	UpdatedAt time.Time       `json:"-" faker:"-"`
	DeletedAt sql.NullTime    `gorm:"index" json:"-" faker:"-"`
	ClientID  uint            `json:"-" faker:"-"`
	Client    client.Client   `json:"client" faker:"-"`
	ProductID uint            `json:"-" faker:"-"`
	Product   product.Product `json:"product" faker:"-"`
	Quantity  uint            `json:"quantity" faker:"-"`
}
