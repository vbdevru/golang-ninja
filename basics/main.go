package main

import "fmt"

func main() {
	defer handlerPanic()
	// defer printMessage()

	// fmt.Println("main()")
	messages := []string{
		"message 1",
		"message 2",
		"message 3",
		"message 4",
	}
    fmt.Println(messages)
	messages[4] = "message 5"

	//fmt.Println(messages)
	//panic("aaaaaaaaaa help")
}
	
func handlerPanic() {
    if r := recover(); r != nil {
		fmt.Println(r)
	}

	fmt.Println("handlerPanic() выполнилась успешно")
}