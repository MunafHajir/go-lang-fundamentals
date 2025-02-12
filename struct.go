package main

import "fmt"

// Define Struct
type Person struct {
	Name string
	Age  int
}

// Method for Struct
func (p Person) greet() {
	fmt.Println("Hello, my name is", p.Name, "and I am", p.Age, "years old.")
}

func main() {
	// Create Object
	p := Person{Name: "Munaf", Age: 25}
	p.greet()
}
