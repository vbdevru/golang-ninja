package main

import (
	"errors"
	"fmt"
)

func main() {
<<<<<<< HEAD
	// var message string
	// message = sayHello("Максим", 21)
	// printMessage(message)
	// printMessage("вызов 1")
	// printMessage("вызов 2")
	// printMessage("вызов 3")
	// mesage, err := enterTheClub(12)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	fmt.Println(prediction("rrgrtg"))
=======
	var message string
	message = sayHello("Максим", 21)
	printMessage(message)
	printMessage("вызов 1")
	printMessage("вызов 2")
	printMessage("вызов 3")
	mesage, err := enterTheClub(12)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(mesage)
>>>>>>> 62091d5c1fa39ef57490d8106510a6018ff690bc
}

func printMessage(message string) {
	fmt.Println(message)
}

func sayHello(name string, age int) string {
	return fmt.Sprintf("Привет, %s! Тебе %d лет", name, age)
}

func enterTheClub(age int) (string, error) {
	if age >= 18 && age <= 45 {
		return "Входи, только аккуратно", nil
	} else if age >= 45 && age < 65 {
		return "Вам точно понравится эта музыка?", nil
	} else if age >= 65 {
		return "Это уже слишком для вас", errors.New("you are too old")
	}

	return "Тебе нет 18-ти", errors.New("you are too young")
	
<<<<<<< HEAD
}

func prediction(dayOfWeek string) (string,error) {
	switch dayOfWeek {
	case "пн":
		return "Хорошего тебе начала недели!", nil
	case "вт":
		return "Прекрасного тебе вторника!", nil
	case "ср":
		return "Замечательной тебе среды!", nil
	case "чт":
		return "Четверг - это маленькая пятница! Только не переборщи!", nil
	default:
		return "невалидный день недели", errors.New("invalid day of the week")
	}
=======
>>>>>>> 62091d5c1fa39ef57490d8106510a6018ff690bc
}