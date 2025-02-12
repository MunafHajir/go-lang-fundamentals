package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup) {
	defer wg.Done() // Mark as done when the function exits
	fmt.Println("Worker started")
	time.Sleep(2 * time.Second) // Simulate work
	fmt.Println("Worker finished")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1) // Add one goroutine to the wait group
	go worker(&wg)

	wg.Wait() // Wait for the goroutine to complete
	fmt.Println("Now executing after goroutine")
}
