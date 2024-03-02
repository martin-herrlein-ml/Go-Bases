package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Cellphone string `json:"cellphone"`
	Address   string `json:"address"`
}

func main() {
	myPerson := Person{"Martin", "Herrlein", "123456", "Pepito 123"}
	personAsJSON, err := json.Marshal(myPerson)

	fmt.Println(string(personAsJSON))
	fmt.Println(err)
}
