package main

import "fmt"

func main() {
	someIDs := []int{1, 2, 3, 4, 534, 645, 74}
	for i, id := range someIDs {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Without using i, because otherise we get an error that i is not used
	for _, id := range someIDs {
		fmt.Printf("ID: %d\n", id)
	}

	sum := 0
	for _, id := range someIDs {
		sum += id
	}
	fmt.Println("Sum is", sum)

	// Something else with range and maps
	employee := map[string]string{
		"name":            "Sultan",
		"title":           "rabotyaga",
		"work experience": "one year",
	}

	for k, v := range employee {
		fmt.Printf("%s is %s\n", k, v)
	}

}
