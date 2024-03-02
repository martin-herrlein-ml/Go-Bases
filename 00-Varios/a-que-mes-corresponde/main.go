package main

import "fmt"

func main() {
	var dayOfMonth int
	fmt.Scan(&dayOfMonth)
	println(month(dayOfMonth))

}

func month(number int) string {
	var month = map[int]string{1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril", 5: "Mayo", 6: "Junio"}
	return month[number]
}
