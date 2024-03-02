package main

import "fmt"

const (
	Addition       = "+"
	Substraction   = "-"
	Multiplication = "*"
	Division       = "/"
)

func calculate(value1, value2 float64, operation string) float64 {
	switch operation {
	case Addition:
		return value1 + value2
	case Substraction:
		return value1 - value2
	case Multiplication:
		return value1 * value2
	case Division:
		if value2 != 0 {
			return value1 / value2
		}
	}
	return 0
}

func main() {
	fmt.Println(calculate(455, 88, Addition))
	fmt.Println(calculate(99, 88, Substraction))
	fmt.Println(calculate(234, 2, Multiplication))
	fmt.Println(calculate(505, 4, Division))

}
