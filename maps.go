package main

import "fmt"

func main() {
	// Map (Key-Value Pairs)
	student := map[string]int{
		"John":  90,
		"Alice": 85,
	}

	// Access value
	fmt.Println("John's Score:", student["John"])

	// Update value
	student["Alice"] = 95

	// Iterate over map
	for key, value := range student {
		fmt.Println(key, ":", value)
	}
}
