package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d is even: is %t\n", 16, even(16)) // 16 is even: is true
	fmt.Printf("%d is odd: is %t\n", 17, odd(17))
	// 17 is odd: is true
	fmt.Printf("%d is odd: is %t\n", -17, odd(-17))
	// 18 is odd: is false
}

func even(nr int) bool {
	fmt.Println(nr)
	if nr == 0 {
		return true
	}
	return odd(RevSign(nr) - 1)
}

func odd(nr int) bool {
	fmt.Println(nr)
	if nr == 0 {
		return false
	}
	return even(RevSign(nr) - 1)
}

func RevSign(nr int) int {
	if nr < 0 {
		return -nr
	}
	return nr
}