package main

import "fmt"

func main() {
	 users := map[string]int{
		"Vasya": 15,
		"Petya": 23,
		"Kostya": 48,
	 }

	 fmt.Println(len(users))
	

	 // users["Serega"] = 21
    //  var  users map[string]int
	//  users = make(map[string]int)

	//  users["Vasya"] = 19

	//  fmt.Println(users)
	 
	//  delete(users, "Vasya")

	//  for key, value := range users {
	// 	fmt.Println(key, value)
	//  }

	//  age, ok := users["Serega"]
	//  if ok {
	// 	fmt.Println("Kostya", age)
	//  } else {
	// 	fmt.Println("нет в списке")
	//  }

	//  age, exists := users["Serega"]
	//  fmt.Println(age, exists)

	//  age2, exists2 := users["Petya"]
	//  fmt.Println(age2, exists2)
}
