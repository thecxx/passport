package proto

type User interface {
	Create() error
	Delete() error
}
