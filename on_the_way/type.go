package main

import (
	"fmt"
)

type Integer int

func (i Integer) add(num int) int {
	return num + int(i)
}

func (i Integer) minus(num int) int {
	return int(i) - num

}

func main() {
	fmt.Println(Integer(2).add(3))
	fmt.Println(Integer(2).minus(3))
}
