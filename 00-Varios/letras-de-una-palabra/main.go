package main

import "fmt"

func main() {
	var word = "HELLO"

	fmt.Println("Cantidad de letras: ", len(word))
	for index, runeValue := range word {
		fmt.Printf("%#U empieza en el byte de posición %d\n", runeValue, index)
	}
}
