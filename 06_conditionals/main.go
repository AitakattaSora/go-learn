package main

import (
	"fmt"
)

func greaterOrLess(x, y int) string {
	if x > y {
		fmt.Printf("%d is greater than %d\n", x, y)
		return "greater"
	} else {
		fmt.Printf("%d is less than %d\n", x, y)
		return "less"
	}
}

func whatColor(color string) {
	switch color {
	case "blue":
		fmt.Println("color is blue")
	case "red":
		fmt.Println("color is red")
	default:
		fmt.Println("I dont know this color")
	}
}

func main() {
	greaterOrLess(1, 4)
	greaterOrLess(24, 4)
}
