package main

import "fmt"

func getSequence() func() int {
	var a int = 1
	var b int = 0
	var i int = 0

	return func() int {
		i++
		if i == 1 {
			return 0
		} else {
			c := a + b
			a = b
			b = c
			return c
		}
	}
}

func main() {
	/* nextNumber is now a function with i as 0 */
	nextNumber := getSequence()

	/* invoke nextNumber to increase i by 1 and return the same */
	for i := 0; i < 80; i++ {
		fmt.Println(nextNumber())
	}
}
