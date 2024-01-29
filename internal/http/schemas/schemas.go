package schemas

type CreateSchema interface {
	ToModel() interface{}
}

type Validatable interface {
	Validate() error
}
