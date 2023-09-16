package interfaces

type Repository[T Model] interface {
	List() ([]T, error)
	Get(id int64) (*T, error)
	Create(c *T) error
	Update(c *T) error
	Delete(c *T) error
}
