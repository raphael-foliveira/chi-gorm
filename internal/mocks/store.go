package mocks

import (
	"errors"

	"github.com/raphael-foliveira/chi-gorm/internal/entities"
)

type store[T entities.Entity] struct {
	Store []T
	Error error
}

func (s *store[T]) List(conds ...interface{}) ([]T, error) {
	return s.Store, s.Error
}

func (s *store[T]) Get(id uint) (*T, error) {
	for _, entity := range s.Store {
		if entity.GetId() == id {
			return &entity, s.Error
		}
	}
	return nil, errors.New("not found")
}

func (s *store[T]) Create(client *T) error {
	s.Store = append(s.Store, *client)
	return s.Error
}

func (s *store[T]) Update(client *T) error {
	for i, c := range s.Store {
		if c.GetId() == (*client).GetId() {
			s.Store[i] = (*client)
			return s.Error
		}
	}
	return errors.New("not found")
}

func (s *store[T]) Delete(client *T) error {
	for i, c := range s.Store {
		if c.GetId() == (*client).GetId() {
			s.Store = append(s.Store[:i], s.Store[i+1:]...)
			return s.Error
		}
	}
	return errors.New("not found")
}

var ClientsStore = &clientsStore{store[entities.Client]{}}

type clientsStore struct {
	store[entities.Client]
}

var ProductsStore = &productsStore{store[entities.Product]{}}

type productsStore struct {
	store[entities.Product]
}

func (cr *productsStore) FindMany(ids []uint) ([]entities.Product, error) {
	products := []entities.Product{}
	for _, id := range ids {
		for _, product := range cr.Store {
			if product.ID == id {
				products = append(products, product)
			}
		}
	}
	return products, cr.Error
}

var OrdersStore = &ordersStore{store[entities.Order]{}}

type ordersStore struct {
	store[entities.Order]
}

func (os *ordersStore) FindManyByClientId(clientId uint) ([]entities.Order, error) {
	orders := []entities.Order{}
	for _, order := range os.Store {
		if order.ClientID == clientId {
			orders = append(orders, order)
		}
	}
	return orders, os.Error
}
