package main

import "fmt"

func main() {
	// How many steps before the final 1 ?
	i := 5
	fmt.Println(i)
	for i != 1 {
		if i % 2 == 0 {
			fmt.Println("i/2\t=", i / 2)
			i = i / 2
		} else {
			//The expression 3i+1 is actually a complex number literal
			fmt.Println("3i+1\t=", 3i + 1)
			i = 3 * i + 1
		}
	}
}