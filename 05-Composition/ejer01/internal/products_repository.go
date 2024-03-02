package internal

type ProductRepository interface {
	GetAll() (products []Product)
	Save(product Product) (products []Product)
	GetById(ID int) (product Product)
}
