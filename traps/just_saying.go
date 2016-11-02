package main

import "fmt"
// If the assertion [doesn't hold], the expression returns (Z, false)
// where Z is the zero value for type T.
// In the case of Rosie, the "type asserted" x variable in the else{...}
// block is a Human with zero value. It shadows the initial Dog value of x.
func main() {
	var x Animal = Dog{"Rosie"}

	if x, ok := x.(Human); ok {
		fmt.Println(ok)
		fmt.Printf("%v doesn't want to be treated like dogs and cats.\n", x.lastName)
	} else {
		fmt.Println(ok)
		fmt.Println(x.Say())
	}
}

type Animal interface {
	Say() string
}

type Dog struct {
	name string
}

func (d Dog) Say() string {
	return fmt.Sprintf("%v barks", d.name)
}

type Cat struct {
	name string
}

func (c Cat) Say() string {
	return fmt.Sprintf("%v meows", c.name)
}

type Human struct {
	firstName string
	lastName  string
}

// Humans are technically animals, and they say things.
func (h Human) Say() string {
	return fmt.Sprintf("%v %v speaks", h.firstName, h.lastName)
}