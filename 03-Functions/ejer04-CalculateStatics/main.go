/*
Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de
calificaciones de los/as estudiantes de un curso.
Requieren calcular los valores mínimo, máximo y promedio de sus calificaciones.

Para eso, se solicita generar una función que indique qué tipo de cálculo se quiere realizar
(mínimo, máximo o promedio) y que devuelva otra función y un mensaje
(en caso que el cálculo no esté definido) que se le pueda pasar una cantidad N de enteros y
devuelva el cálculo que se indicó en la función anterior.
*/

package main

import "fmt"

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

type TypeOperation func(...int) (int, string)

func minFunc()

/*
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

*/

func operation(typeOperation string) {
	switch typeOperation {
	case minimum:
		return minFunc()
	case average:
		return averageFunc()
	case maximum:
		return maxFunc()
	default:
		return fmt.Errorf("No está definido el tipo de cálculo")

	}
}
