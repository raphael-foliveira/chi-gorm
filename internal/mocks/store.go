package mocks

import (
	"errors"

	"github.com/go-faker/faker/v4"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
)

var ProductsStore = newProductsStore()
var ClientsStore = newClientsStore()
var OrdersStore = newOrdersStore()
var UsersStore = newUsersStore()

type store[T entities.Entity] struct {
	Store []T
	Error error
}

func newStore[T entities.Entity]() *store[T] {
	return &store[T]{}
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

func (s *store[T]) Clear() {
	s.Store = []T{}
}

func (s *store[T]) Populate() {
	s.Store = []T{}
	faker.FakeData(&s.Store)
}

type clientsStore struct {
	*store[entities.Client]
}

func newClientsStore() *clientsStore {
	return &clientsStore{newStore[entities.Client]()}
}

type productsStore struct {
	*store[entities.Product]
}

func newProductsStore() *productsStore {
	return &productsStore{newStore[entities.Product]()}
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

type ordersStore struct {
	*store[entities.Order]
}

func newOrdersStore() *ordersStore {
	return &ordersStore{newStore[entities.Order]()}
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

type usersStore struct {
	*store[entities.User]
}

func newUsersStore() *usersStore {
	return &usersStore{newStore[entities.User]()}
}

func (u *usersStore) FindOneByEmail(email string) (*entities.User, error) {
	for _, user := range u.Store {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}
