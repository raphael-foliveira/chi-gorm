package schemas

type CreateSchema interface {
	ToModel() interface{}
}
