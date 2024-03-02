package main

import (
	"ejer02/composition/internal"
	"fmt"
)

func main() {

	employe01 := internal.Employee{
		ID:       0004,
		Position: "Developer",
		Person: internal.Person{
			Name:        "Martin Herrlein",
			ID:          4689,
			DateOfBirth: "22/02/1992",
		},
	}

	fmt.Println(employe01.PrintEmployee())
}
