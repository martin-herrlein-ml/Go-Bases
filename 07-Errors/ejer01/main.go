package main

import "fmt"

type salaryError struct {
	msg string
}

func (e *salaryError) Error() string {
	return e.msg
}

func main() {
	salary := 190000

	if salary < 150000 {
		err := salaryError{"Error: the salary entered does not reach the taxable minimum"}
		fmt.Println(err.Error())
	} else {
		fmt.Println("Must pay tax")
	}

}
