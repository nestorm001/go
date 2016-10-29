package main

import "fmt"

func main() {
	nephews := []string{"Huey", "Dewey"}
	nephews = append(nephews, "Louie")

	for nephew := range nephews {
		fmt.Println("Hello", nephew)
	}

	fmt.Println("===real way===")
	for _, nephew := range nephews {
		fmt.Println("Hello", nephew)
	}
}