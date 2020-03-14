package main

import (
	"fmt"
)

func greeting(name string) string {
	return "Hello " + name
}

func getSumm(num1 int, num2 int) int {
	fmt.Println(num1 + num2)
	return num1 + num2
}
func main() {
	fmt.Println(greeting("Sultan"))
	getSumm(13, 14)
}
