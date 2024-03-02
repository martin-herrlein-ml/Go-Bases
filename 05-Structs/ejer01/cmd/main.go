package main

import (
	"ejer/composition/internal"
	"ejer/composition/internal/repository"
	"ejer/composition/internal/service"
	"fmt"
)

func main() {
	var repo internal.ProductRepository
	db := []internal.Product{
		{
			ID: 0001,
			ProductAttributes: internal.ProductAttributes{
				Category:    "Hardware",
				Name:        "Notebook",
				Price:       499,
				Description: "Asus VivBook i7",
			},
		},
		{
			ID: 0025,
			ProductAttributes: internal.ProductAttributes{
				Category:    "Hardware",
				Name:        "Memoria RAM",
				Price:       499,
				Description: "16GB DDR5 Kingston",
			},
		},
		{
			ID: 0432,
			ProductAttributes: internal.ProductAttributes{
				Category:    "Hardware",
				Name:        "SSD 1TB",
				Price:       499,
				Description: "Pci - E M.2",
			},
		},
		{
			ID: 0011,
			ProductAttributes: internal.ProductAttributes{
				Category:    "Hardware",
				Name:        "Monitor",
				Price:       499,
				Description: "Samsung 24' FHD",
			},
		},
	}
	repo = repository.NewProductsSlice(db)
	serv := service.NewProductDefault(repo)

	addProduct := internal.Product{
		ID: 15,
		ProductAttributes: internal.ProductAttributes{
			Category:    "Hardware",
			Name:        "Cooler",
			Price:       499,
			Description: "Coolermaster turbo",
		},
	}

	serv.Save(addProduct)

	products, err := serv.GetAll()
	fmt.Println(len(products))
	if err != "" {
		println(err)
		return
	}

	for _, element := range products {
		fmt.Println(element)
	}
	/*
		addProduct, err = serv.GetById(0011)
		if err != "" {
			println(err)
			return
		}
		fmt.Println("Obtener por ID 0011", addProduct)

	*/

}
