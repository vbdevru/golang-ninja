package main

import (
	"fmt"
)

type Age int 

func (a Age) isAdult() bool {
	return a >= 18
}

type User struct {
	name   string
	age    Age
	sex    string
	weight int
	height int
}

func (u *User) setName(name string) {
	u.name = name
}

func (u User) getName() string {
	return u.name
}

type DumbDatabase struct {
	m map[string]string
}

func NewDumbDatabase() *DumbDatabase {
	return &DumbDatabase{
		m: make(map[string]string),
	}
}

func NewUser(name, sex string, age, weight, height int) User {
	return User{
		name: name,
		sex: sex,
		age: Age(age),
		weight: weight,
		height: height,
	}
}

func main() {
	user1 := NewUser("Vasya", "Male", 12, 75, 186)
	user2 := User{"Petya", 31, "Male", 84, 197}

	// user1.printUserInfo("Kostya")
	// user2.printUserInfo("Serega")

	user1.setName("Serega")

	fmt.Println(user1.age.isAdult())

	fmt.Println(user1.getName())
	fmt.Println(user2.getName())

	// printUserInfo(user1)
	// printUserInfo(user2)

    // fmt.Println(user1.name, user1.age)
	// fmt.Println(user2.name)

	// user := struct {
	// 	name   string
	// 	age    int
	// 	sex    string
	// 	weight int
	// 	height int
	// }{"Vasya", 23, "Male", 75, 185}

	// fmt.Printf("%+v\n", user)
	// fmt.Printf("%+v\n", user1)
	// fmt.Printf("%+v\n", user2)
}

// func printUserInfo(user User) {
// 	fmt.Println(user.name, user.age, user.sex, user.weight, user.height)
// }