package main

import (
	"fmt"
)

func main() {
	callback(2, 1, Add)
	callback(2, 1, Divide)
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func Divide(a, b int) {
	fmt.Printf("The result of %d divide %d is: %d\n", a, b, 1.0 *a/b)
}

func callback(x int, y int, f func(int, int)) {
	f(x, y)
}