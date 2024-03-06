package main

import (
	"errors"
	"fmt"
)

const (
	limit = 10000
)

type CustomError struct {
	message string
}

func (e *CustomError) Error() string {
	return e.message //
}

func main() {
	salary := 210

	_, err := calculator(salary)

	if err != nil {
		if errors.Is(err, &CustomError{}) {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Must pay tax")

}

func calculator(salary int) (result bool, err error) {
	if salary < limit {
		return false, &CustomError{"Error: salary is less than or equal to 10000"}
	}
	return
}
