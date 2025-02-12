package main

import "fmt"

func main() {
	num := 10

	// If-Else
	if num%2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}

	// Loops (For Loop)
	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration:", i)
	}
}
