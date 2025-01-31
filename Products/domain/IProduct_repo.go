package domain

type Iproduct interface {
	Save(product *Product) error
	GetAll() ([]Product, error)
	Edit(name string, price float32, id int32) error
	Delete(id int32) error
}