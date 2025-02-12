package main

import "fmt"

// Function with parameters and return value
func add(a int, b int) int {
	return a + b
}

func main() {
	sum := add(5, 7)
	fmt.Println("Sum:", sum)
}
