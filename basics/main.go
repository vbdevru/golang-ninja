package main

import "fmt"

func main() {
	// messages := []string{
	// 	"message 1",
	// 	"message 2",
	// 	"message 3",
	// 	"message 4",
	// }

	// for i := 0; i < len(messages); i++ {
	// 	fmt.Println(messages[i])
	// }

	// for _, message := range messages {
	// 	fmt.Println(message)
	// }


    counter := 0
	for {
		if counter == 100 {
			break
		}

		counter++
		fmt.Println(counter)
	}
}
	