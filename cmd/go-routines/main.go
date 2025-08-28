package main

import (
	"fmt"
	"sync"
	"time"
)

func say(message string, wg *sync.WaitGroup) {
	// postpone wg.Done() to ensure the WaitGroup is notified
	// that this goroutine has finished, even if a panic occurs.
	defer wg.Done()

	for i := 0; i < 3; i++ {
		fmt.Println(message)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// --- The Problem: The Program Terminates Too Soon ---
	// If we simply called `go say("Hello")`, the main() function would not wait
	// for the goroutine to complete, and the program would terminate immediately.
	// To solve this, we use a `sync.WaitGroup`.

	// A WaitGroup waits for a collection of goroutines to complete.
	var wg sync.WaitGroup

	// We add 2 to the WaitGroup counter, because we are going to start 2 goroutines.
	wg.Add(2)

	// To start a goroutine, just use the `go` keyword
	go say("Hello", &wg)
	go say("World", &wg)

	// wg.Wait() blocks the execution of the main() function until the WaitGroup's counter
	// reaches zero.
	// The counter is decremented by calling `wg.Done()` within each goroutine.
	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()

	fmt.Println("All goroutines finished... Exiting.")

}
