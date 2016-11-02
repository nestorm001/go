package main

import "fmt"

// Map indexing with non-existent index doesn't panic, it yields the zero value of the element type.
// Use the two variables assignment v, ok := m["four"] to check if requested index was actually
// in the map or not. It is called the "comma ok" idiom.
func main() {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(m["four"])

	v, ok := m["four"]
	fmt.Println(v)
	fmt.Println(ok)
	v, ok = m["three"]
	fmt.Println(v)
	fmt.Println(ok)
}