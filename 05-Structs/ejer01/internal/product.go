package internal

// Product is a struct that represent the attributes of a products
type ProductAttributes struct {
	Name        string
	Price       float64
	Description string
	Category    string
}

type Product struct {
	ID int
	ProductAttributes
}
