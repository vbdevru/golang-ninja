package main

import (
	"fmt"
	"math"
)

type Shape interface {
	ShapeWithArea
	ShapeWithPerimetr
}

type ShapeWithArea interface {
	Area() float32
}

type ShapeWithPerimetr interface {
	Perimeter() float32
}

type Square struct {
	sideLenght float32
}

func (s Square) Area() float32 {
	return s.sideLenght * s.sideLenght
}

func (s Square) Perimeter() float32 {
	return s.sideLenght * 4
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func (c Circle) Perimeter() float32 {
	return 2 * c.radius * c.radius
}

func main() {
	square := Square{5}
	circle := Circle{8}

	printShapeArea(square)
	printShapeArea(circle)

	// printInterface(square)
	// printInterface(circle)
	printInterface(99)
	printInterface("str")
}

func printShapeArea(shape Shape) {
	fmt.Println(shape.Area())
	fmt.Println(shape.Perimeter())
}

func printInterface(i interface{}) {
	// switch value := i.(type) {
	// case int:
	// 	fmt.Println("int", value)
	// case bool:
	// 	fmt.Println("bool", value)
	// default:
	// 	fmt.Printf("unknown type", value)
	// }
	// fmt.Printf("%+v\n", i)

	str, ok := i.(string)
	if !ok {
		fmt.Println("interface is not string")
		return
	}
	fmt.Println(len(str))
}