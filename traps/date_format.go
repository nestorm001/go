package main

import "fmt"
import "time"

func main() {
	// wrong format
	t, err := time.Parse("2006-01-01", "2006-01-02")
	fmt.Println(t)
	fmt.Println(err)
}