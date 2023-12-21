package main

import (
	"fmt"
	"golang-ninja/basic/shape"
)

func main() {
	square := shape.NewSquare(5)
	fmt.Println(square)
}