package main

import "fmt"

//The short declaration operator := declares a new variable,
// and is capable of shadowing a previously declared variable.
// Outer variable i is never modified and the loop keeps spinning.
// New variables with same name i are declared (by mistake) inside the if/else blocks.
func main() {
	// If this loops forever, I might have broken the Collatz conjecture!!
	i := 1859
	fmt.Println(i)
	for i != 1 {
		if i % 2 == 0 {
			//i := i / 2
			i = i / 2
			fmt.Println("i/2\t=", i)
		} else {
			//i := 3 * i + 1
			i = 3 * i + 1
			fmt.Println("3i + 1\t=", i)
		}
	}
}