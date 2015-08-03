package main

import (
	"fmt"
	"time"
)

func getSequence() (func() int) {
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
	start := time.Now()
	/* invoke nextNumber to increase i by 1 and return the same */
	for i := 0; i < 80; i++ {
		fmt.Println(nextNumber())
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}
