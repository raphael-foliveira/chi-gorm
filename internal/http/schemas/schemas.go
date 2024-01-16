package schemas

type CreateSchema interface {
	ToModel() interface{}
}

type ValidateableSchema interface {
	Validate() error
}
