package main

import "fmt"

func calcImp(salary float64) (tax float64) {

	if salary >= 50000 {
		tax = salary * 0.17
	}
	if salary >= 150000 {
		tax += salary * 0.9
	}
	return tax
}

func main() {
	fmt.Println(calcImp(49000))
}
