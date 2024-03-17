package mocks

import (
	"errors"

	"github.com/go-faker/faker/v4"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
)

type repo[T entities.Entity] struct {
	Store []T
	Error error
}

func (s *repo[T]) List(conds ...interface{}) ([]T, error) {
	return s.Store, s.Error
}

func (s *repo[T]) Get(id uint) (*T, error) {
	for _, entity := range s.Store {
		if entity.GetId() == id {
			return &entity, s.Error
		}
	}
	return nil, errors.New("not found")
}

func (s *repo[T]) Create(client *T) error {
	s.Store = append(s.Store, *client)
	return s.Error
}

func (s *repo[T]) Update(client *T) error {
	for i, c := range s.Store {
		if c.GetId() == (*client).GetId() {
			s.Store[i] = (*client)
			return s.Error
		}
	}
	return errors.New("not found")
}

func (s *repo[T]) Delete(client *T) error {
	for i, c := range s.Store {
		if c.GetId() == (*client).GetId() {
			s.Store = append(s.Store[:i], s.Store[i+1:]...)
			return s.Error
		}
	}
	return errors.New("not found")
}

var ClientsRepository = &clientsRepository{repo[entities.Client]{}}

type clientsRepository struct {
	repo[entities.Client]
}

var ProductsRepository = &productsRepository{repo[entities.Product]{}}

type productsRepository struct {
	repo[entities.Product]
}

func (cr *productsRepository) FindMany(ids []uint) ([]entities.Product, error) {
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

var OrdersRepository = &ordersRepository{repo[entities.Order]{}}

type ordersRepository struct {
	repo[entities.Order]
}

func (os *ordersRepository) FindManyByClientId(clientId uint) ([]entities.Order, error) {
	orders := []entities.Order{}
	for _, order := range os.Store {
		if order.ClientID == clientId {
			orders = append(orders, order)
		}
	}
	return orders, os.Error
}

func Populate() {
	addClientsAndProducts(10)
	for i := range ClientsRepository.Store {
		addOrderToClient(&ClientsRepository.Store[i])
	}
}

func addClientsAndProducts(q int) {
	for i := 0; i < q; i++ {
		var client entities.Client
		var product entities.Product
		faker.FakeData(&client)
		faker.FakeData(&product)
		ProductsRepository.Store = append(ProductsRepository.Store, product)
		client.ID = uint(i + 1)
		for j := 0; j < 10; j++ {
			addOrderToClient(&client)
		}
		ClientsRepository.Store = append(ClientsRepository.Store, client)
	}
}

func addOrderToClient(client *entities.Client) {
	var product entities.Product
	faker.FakeData(&product)
	ProductsRepository.Store = append(ProductsRepository.Store, product)
	var order entities.Order
	faker.FakeData(&order)
	order.ClientID = client.ID
	order.ProductID = product.ID
	OrdersRepository.Store = append(OrdersRepository.Store, order)
}

var Repositories = &repository.Repositories{
	ClientsRepository:  ClientsRepository,
	OrdersRepository:   OrdersRepository,
	ProductsRepository: ProductsRepository,
}

func ClearRepositories() {
	ClientsRepository.Store = []entities.Client{}
	OrdersRepository.Store = []entities.Order{}
	ProductsRepository.Store = []entities.Product{}
}
