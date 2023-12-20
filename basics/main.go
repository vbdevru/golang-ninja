package main

import "fmt"

func main() {
	number := 5
	var p *int

	p = &number
	// message := "Скоро я стану ниндзя!"

	fmt.Println(p)
	fmt.Println(&number)

	*p = 10

	fmt.Println(number)

	//changeMessage(message)

	//fmt.Println(&message)
}

func changeMessage(message string) {
	// *message += " (из функции printMessage())"
	fmt.Println(&message)
}