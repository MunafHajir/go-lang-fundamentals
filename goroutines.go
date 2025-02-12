package main

// Why Use Goroutines Instead of Threads?
// --------------------------------------
//
// 1. **Creation Time**
//    - Goroutines are created very fast compared to OS threads.
//    - Goroutines: Fast 🚀
//    - OS Threads: Slow 🐢
//
// 2. **Memory Usage**
//    - Goroutines consume very little memory (~2KB per goroutine).
//    - OS Threads require significantly more (~1MB per thread).
//
// 3. **Scheduling**
//    - Goroutines are scheduled by the Go runtime (efficient and automatic).
//    - OS Threads are scheduled by the operating system (heavier and slower).
//
// 4. **Scalability**
//    - Go can efficiently manage millions of goroutines without much overhead.
//    - OS has a limit on the number of threads that can be created.

import (
	"fmt"
	"time"
)

// Function to execute concurrently
func printMessage() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Hello from Goroutine:", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	// Start a new Goroutine
	go printMessage()

	// Main function execution
	fmt.Println("Hello from main function")

	// Wait to see Goroutine output (Otherwise, the program may exit)
	time.Sleep(3 * time.Second)
}
