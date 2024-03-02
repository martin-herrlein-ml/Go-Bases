/*
Ejercicio 2 - Calcular promedio


Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones.
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio.
No se pueden introducir notas negativas.
*/

package main

import "fmt"

func calculateHalf(values ...int) (float32, error) {
	sum := 0
	for _, nota := range values {
		if nota < 0 {
			return 0, fmt.Errorf("No se permiten notas negativas")
		}
		sum += nota
	}

	return float32(sum / len(values)), nil
}

func main() {
	notes := []int{8, 7, 6, -5, 4, 10, 9, 9, 8, 8}
	half, err := calculateHalf(notes...)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Promedio: ", half)
	}

}
