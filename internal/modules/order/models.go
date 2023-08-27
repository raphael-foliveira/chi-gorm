package order

import (
	"database/sql"
	"time"

	"github.com/raphael-foliveira/chi-gorm/internal/modules/product"
)

type Order struct {
	ID        uint            `json:"id" gorm:"primarykey"`
	CreatedAt time.Time       `json:"-"`
	UpdatedAt time.Time       `json:"-"`
	DeletedAt sql.NullTime    `json:"-"`
	ClientID  uint            `json:"clientId"`
	ProductID uint            `json:"productId"`
	Product   product.Product `json:"product"`
	Quantity  uint            `json:"quantity"`
}
