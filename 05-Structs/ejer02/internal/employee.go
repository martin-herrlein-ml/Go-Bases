package internal

import "fmt"

type Employee struct {
	ID       int
	Position string
	Person
}

func (e Employee) PrintEmployee() (info string) {
	info = fmt.Sprintf(
		"ID Employee: %d\nPosition: %s\nPerson: %s\n",
		e.ID,
		e.Position,
		e.PrintPerson(),
	)
	return
}
