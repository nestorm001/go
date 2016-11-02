package main

import "fmt"

// The method receiver, just like the arguments, is always passed by value.
// This means that when the method reflect is called, the Point is actually
// copied, then the copy is transformed and thrown away.

type Point struct{ x, y int }

// Diagonal symmetry: just swap x and y
func (p *Point) reflect(){
	p.x, p.y = p.y, p.x
}

func main() {
	s := []Point{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	s[0].reflect()
	s[1].reflect()

	fmt.Println(s)
}