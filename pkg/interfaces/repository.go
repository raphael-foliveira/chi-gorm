package interfaces

type Repository[T interface{}] interface {
	List() ([]T, error)
	Get(id uint64) (T, error)
	Create(c *T) error
	Update(c *T) error
	Delete(c *T) error
}
