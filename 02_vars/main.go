package main

import (
	"fmt"
)

func main() {
	// Using var
	var name = "John"
	var age = 22
	const isCool = true

	// Shorthand
	length := "132"

	// Multiple declarations
	width, height := 12, 14

	fmt.Println(name, age)
	fmt.Printf("%T\n", "age")
	fmt.Printf("%T\n", isCool)
	fmt.Println(width, height, length)
}
