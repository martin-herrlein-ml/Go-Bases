package internal

import "fmt"

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

func (p Person) PrintPerson() (info string) {
	info = fmt.Sprintf(
		"IDperson: %d\nName: %s\nDateOfBirth: %s\n",
		p.ID,
		p.Name,
		p.DateOfBirth,
	)
	return
}
