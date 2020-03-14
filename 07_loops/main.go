package main

import (
	"fmt"
)

func main() {
	// Long way
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("")

	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)
	}

	// Fizz Buzz
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Printf("%d: Fizz\n", i)
		} else if i%5 == 0 {
			fmt.Printf("%d: Buzz\n", i)
		} else if i%3 == 0 && i%5 == 0 { // or i%15
			fmt.Printf("%d: FizzBuzz\n", i)
		} else {
			fmt.Println(i)
		}
	}
}
