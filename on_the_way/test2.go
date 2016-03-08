package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	const num = 1e8
	a := [num]int{}
	i := 0
	for i < num {
		a[i] = i
		i++
	}

	c := make(chan int)
	go sum(a[:len(a)/4], c)
	go sum(a[len(a)/4:len(a)/2], c)
	go sum(a[len(a)/2:len(a)*3/4], c)
	go sum(a[len(a)*3/4:], c)
	v1, v2, v3, v4 := <-c, <-c, <-c, <-c // receive from c

	fmt.Println(v1 + v2 + v3 + v4)
}
