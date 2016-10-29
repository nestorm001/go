package main

import (
	"fmt"
	"sync"
)

func main() {
	twoNephews := []string{"Huey", "Dewey"}
	threeNephews := append(twoNephews, "Louie")

	var wg sync.WaitGroup

	for i := range threeNephews {
		// increase the wait group counter
		wg.Add(1)
		nephew := threeNephews[i]
		go func() {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			fmt.Println("Hello", nephew)
		}()
	}
	// wait for all wait group finish
	wg.Wait()
}