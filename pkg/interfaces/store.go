package interfaces

type Store[T interface{}] interface {
	List() ([]T, error)
	Get(id int64) (*T, error)
	Create(c *T) error
	Update(c *T) error
	Delete(c *T) error
}
