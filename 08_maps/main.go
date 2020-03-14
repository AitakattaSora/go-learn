package main

import "fmt"

func main() {
	// Define map
	employee := make(map[string]string)

	// Assing values
	employee["name"] = "Bob"
	employee["surname"] = "Ross"
	employee["title"] = "junior QA"

	fmt.Println(employee)

	// Define and assign
	streamers := map[string]string{
		"name":    "Tyler",
		"surname": "Fking one",
		"channel": "loltyler1",
		"title":   "retard",
	}

	fmt.Println(streamers)
	fmt.Println(len(streamers))
	fmt.Println(streamers["name"])
}
