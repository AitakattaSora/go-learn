package main

import (
	"fmt"
	"math"

	"github.com/sultanmomynov/go-learn/03_packages/strutils"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))

	fmt.Println(strutils.Reverse("hello"))
}
