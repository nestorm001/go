package main

import "fmt"
import "sync"
/**
* This is a classic problem when closure elements
* are accessed by concurrent processes. Here the value of the
* iteration variable nephew changes at each step
* so the most probable output is Louie being greeted three times. You
* must capture the current value of nephew
*
* either by declaring a "copying" variable inside the for block Try it,
* or by passing nephew as a function argument Try it.
*
*/
func main() {
	twoNephews := []string{"Huey", "Dewey"}
	threeNephews := append(twoNephews, "Louie")

	var wg sync.WaitGroup

	for _, nephew := range threeNephews {
		wg.Add(1)
		go func() {
			fmt.Println("Hello", nephew)
			wg.Done()
		}()
	}

	// Wait for all greetings to complete.
	wg.Wait()
}