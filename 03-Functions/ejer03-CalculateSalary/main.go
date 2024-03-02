/*
Ejercicio 3 - Calcular salario

Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas
trabajadas por mes y la categoría.

Categoría C, su salario es de $1.000 por hora.
Categoría B, su salario es de $1.500 por hora, más un 20 % de su salario mensual.
Categoría A, su salario es de $3.000 por hora, más un 50 % de su salario mensual.
Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes,
la categoría y que devuelva su salario.
*/
package main

import "fmt"

const (
	categoryA   = "A"
	categoryB   = "B"
	categoryC   = "C"
	salaryHourA = 3000
	salaryHourB = 1500
	salaryHOurC = 1000
)

func calculateSalary(minWork int, category string) (float64, error) {
	var baseSalary float64
	var percentageAdd float64

	switch category {
	case categoryA:
		baseSalary = salaryHourA
		percentageAdd = 0.5
	case categoryB:
		baseSalary = salaryHourB
		percentageAdd = 0.2
	case categoryC:
		baseSalary = salaryHOurC
		percentageAdd = 0
	default:
		return 0, fmt.Errorf("No se encontró la categoría")
	}

	monthSalary := float64(minWork) / 60 * baseSalary * float64(30)
	totalSalary := monthSalary + (monthSalary * percentageAdd)

	return totalSalary, nil
}

func main() {
	minWork := 1800
	category := "B"

	salary, err := calculateSalary(minWork, category)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("El salario es: %.2f\n", salary)
	}

}
