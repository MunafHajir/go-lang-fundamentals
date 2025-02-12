package main

import "fmt"

func main() {
	// Arrays (Fixed Size)
	var arr = [3]int{1, 2, 3}

	// Slices (Dynamic Size)
	slice := []int{4, 5, 6}

	// Append to Slice
	slice = append(slice, 7)

	fmt.Println("Array:", arr)
	fmt.Println("Slice:", slice)
}
