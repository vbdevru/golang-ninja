package main

import (
	"errors"
	"fmt"
)

func main() {
	// messages := [5]string{"1", "2", "3"}

	// messages := []string{"1", "2", "3"}

	// var messages []string

	// printMessages(messages)

	// messages := make([]string, 100)

	// messages[0] = "1"

	// messages = append(messages, "6")
	// messages = append(messages, "7")
	// messages = append(messages, "8")
	// messages = append(messages, "9")
	// messages = append(messages, "10")

	// fmt.Println(messages)
    // fmt.Println(len(messages))
	// fmt.Println(cap(messages))
	// messages = append(messages, "101")
	// fmt.Println(len(messages))
	// fmt.Println(cap(messages))

	matrix := make([][]int, 10)

	counter := 0
	for x := 0; x < 10; x++ {
		matrix[x] = make([]int, 10)
		for y := 0; y < 10; y++ {
			counter++
			matrix[x][y] = counter
		}
		fmt.Println(matrix[x])
	} 

}

func printMessages(messages []string) error {
	if len(messages) == 0 {
		return errors.New("empty array")
	}

	messages[1] = "5"

	fmt.Println(messages)

	return nil
}
