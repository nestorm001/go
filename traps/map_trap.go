package main

import "fmt"

func main() {
	//There is no guarantee regarding the order of iteration in a hash map.
	// The range order for a map is not the order in which the elements were inserted.

	//A very nice feature of maps is that the runtime randomizes map iteration order.
	// This alerts early a beginner who would by mistake rely on iteration order
	// being the insertion order, or iteration order being stable between
	//executions of the loop.
	m := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	}
	for name, value := range m {
		fmt.Println(name, "is english for number", value)
	}
}