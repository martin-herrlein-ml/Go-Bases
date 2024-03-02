package service

import (
	"ejer/composition/internal"
)

var (
	ErrMsgNoProduct      = "No hay productos"
	ErrMsgInsertProduct  = "Debe ingresar un producto"
	ErrMsgSendID         = "Debe especificar un ID"
	ErrMsgNoExistProduct = "No existe producto con el ID ingresado"
)

func NewProductDefault(repository internal.ProductRepository) ProductDefault {
	return ProductDefault{repository}
}

type ProductDefault struct {
	repository internal.ProductRepository
}

func (p ProductDefault) GetAll() (product []internal.Product, err string) {
	product = p.repository.GetAll()

	if len(product) == 0 {
		err = ErrMsgNoProduct
		return
	}

	return product, ""
}

func (p ProductDefault) Save(prod internal.Product) (product []internal.Product, err string) {

	if prod.ID == 0 {
		err = ErrMsgNoProduct
		return
	}
	product = p.repository.Save(prod)

	if len(product) == 0 {
		err = ErrMsgNoProduct
		return
	}

	return product, ""
}

func (p ProductDefault) GetById(ID int) (product internal.Product, err string) {
	if ID == 0 {
		err = ErrMsgSendID
		return
	}
	product = p.repository.GetById(ID)
	if product.ID == 0 {
		err = ErrMsgNoExistProduct
		return
	}
	return product, ""
}
