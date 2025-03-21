//go:build unit

package mocks

import (
	"errors"

	"github.com/go-faker/faker/v4"
	"github.com/raphael-foliveira/chi-gorm/internal/domain"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
	"github.com/stretchr/testify/mock"
)

type Repo[T interface{}] struct {
	mock.Mock
}

func (r *Repo[T]) ExpectSuccess() {
	r.ExpectedCalls = nil
	r.On("Create", mock.Anything).Return(nil)
	r.On("Update", mock.Anything).Return(nil)
	r.On("Delete", mock.Anything).Return(nil)
}

func (r *Repo[T]) ExpectError() {
	r.ExpectedCalls = nil
	r.On("List").Return([]T{}, errMock)
	r.On("Get", mock.Anything).Return(nil, errMock)
	r.On("Create", mock.Anything).Return(errMock)
	r.On("Update", mock.Anything).Return(errMock)
	r.On("Delete", mock.Anything).Return(errMock)
}

var errMock = errors.New("mock error")

func (s *Repo[T]) List(conds ...interface{}) ([]T, error) {
	c := s.Called(conds...)
	return c.Get(0).([]T), c.Error(1)
}

func (s *Repo[T]) Get(id uint) (*T, error) {
	c := s.Called(id)
	result, _ := c.Get(0).(*T)
	return result, c.Error(1)
}

func (s *Repo[T]) Create(client *T) error {
	c := s.Called(client)
	return c.Error(0)
}

func (s *Repo[T]) Update(client *T) error {
	c := s.Called(client)
	return c.Error(0)
}

func (s *Repo[T]) Delete(client *T) error {
	c := s.Called(client)
	return c.Error(0)
}

type clientsRepository struct {
	Repo[domain.Client]
}

func (cr *clientsRepository) ExpectSuccess() {
	cr.Repo.ExpectSuccess()
	cr.On("List").Return(ClientsStub, nil)
	cr.On("Get", mock.Anything).Return(&ClientsStub[0], nil)
}

type productsRepository struct {
	Repo[domain.Product]
}

func (cr *productsRepository) FindMany(ids []uint) ([]domain.Product, error) {
	c := cr.Called(ids)
	return c.Get(0).([]domain.Product), c.Error(1)
}

func (pr *productsRepository) ExpectSuccess() {
	pr.Repo.ExpectSuccess()
	pr.On("List", mock.Anything).Return(ProductsStub, nil)
	pr.On("Get", mock.Anything).Return(&ProductsStub[0], nil)
	pr.On("FindMany", mock.Anything).Return(ProductsStub, nil)
}

func (pr *productsRepository) ExpectError() {
	pr.Repo.ExpectError()
	pr.On("FindMany", mock.Anything).Return(nil, errMock)
}

type ordersRepository struct {
	Repo[domain.Order]
}

func (os *ordersRepository) FindManyByClientId(clientId uint) ([]domain.Order, error) {
	c := os.Called(clientId)
	return c.Get(0).([]domain.Order), c.Error(1)
}

func (or *ordersRepository) ExpectSuccess() {
	or.Repo.ExpectSuccess()
	or.On("List").Return(OrdersStub, nil)
	or.On("Get", mock.Anything).Return(&OrdersStub[0], nil)
	or.On("FindManyByClientId", mock.Anything).Return(OrdersStub, nil)
}

func (or *ordersRepository) ExpectError() {
	or.Repo.ExpectError()
	or.On("FindManyByClientId", mock.Anything).Return(nil, errMock)
}

func init() {
	faker.FakeData(&ClientsStub)
	faker.FakeData(&ProductsStub)
	faker.FakeData(&OrdersStub)
	for i := range ClientsStub {
		ClientsStub[i].Orders = []domain.Order{}
	}
}

func Repositories() {
	OrdersRepository.ExpectSuccess()
	ProductsRepository.ExpectSuccess()
	ClientsRepository.ExpectSuccess()
	repository.Clients = ClientsRepository
	repository.Products = ProductsRepository
	repository.Orders = OrdersRepository
}

func ClearRepositoryMocks() {
	ClientsRepository.ExpectedCalls = nil
	ProductsRepository.ExpectedCalls = nil
	OrdersRepository.ExpectedCalls = nil
}
