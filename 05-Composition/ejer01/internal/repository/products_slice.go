package repository

import "ejer/composition/internal"

func NewProductsSlice(products []internal.Product) ProductsSlice {
	return ProductsSlice{products}
}

type ProductsSlice struct {
	products []internal.Product
}

func (p ProductsSlice) GetAll() (product []internal.Product) {
	product = make([]internal.Product, len(p.products))
	copy(product, p.products)
	return
}

func (p ProductsSlice) Save(product internal.Product) (products []internal.Product) {
	p.products = append(p.products, product)
	return p.GetAll()
}

func (p ProductsSlice) GetById(ID int) (product internal.Product) {
	for _, element := range p.products {
		if element.ID == ID {
			return element
		}
	}
	return internal.Product{}
}
