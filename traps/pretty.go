package main

import "fmt"

const PLUS = '+'
const IMG_UNIT = 'i'

// int plus char

func main() {
	real := 4
	img := 3
	complex := real + PLUS + img + IMG_UNIT
	fmt.Println(complex)
}