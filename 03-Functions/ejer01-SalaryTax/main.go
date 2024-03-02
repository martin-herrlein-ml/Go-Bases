/*

Ejercicio 1 - Impuestos de salario
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar
el sueldo, para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.

Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo y si gana
más de $150.000 se le descontará además un 10 % (27% en total).

*/

package main

import "fmt"

const (
	percentageBase = 0.17
	percentageTop  = 0.10
	salaryBase     = 50000
	salaryTop      = 150000
)

func CalcTax(salary float64) (tax float64) {

	if salary >= salaryBase {
		tax = salary * percentageBase
	}
	if salary >= salaryTop {
		tax += salary * percentageTop
	}
	return tax
}

func main() {
	var salary float64 = 40000
	fmt.Println("Salario:", salary, "Impuesto:", CalcTax(salary))

	salary = 160000
	fmt.Println("Salario:", salary, "Impuesto:", CalcTax(salary))

	salary = 55000
	fmt.Println("Salario:", salary, "Impuesto:", CalcTax(salary))
}
