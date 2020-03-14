package main

import (
	"fmt"
)

func main() {
	// Arrays have to:
	// have a fixed length
	// be a fixed type
	var fruits [2]string // [2] - length
	fruits[0] = "Apple"
	fruits[1] = "Orange"

	veges := [2]string{"Potato", "Tomato"}

	fmt.Println(fruits, veges)

	// Slices - Arrays without fixed length
	vegesSlices := []string{"Cucumber", "Something else", "Something else else"}

	fmt.Println(vegesSlices)
}
