package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	// Here are some calculations with bigInts:
	im := big.NewInt(math.MaxInt64)
	in := im
	io := big.NewInt(1956)
	ip := big.NewInt(1)
	ip.Mul(im, in).Add(ip, im).Div(ip, io)
	fmt.Printf("Big Int: %v\n", ip)
	// Here are some calculations with bigInts:
	rm := big.NewRat(math.MaxInt64, 1956)
	rn := big.NewRat(-1956, math.MaxInt64)
	ro := big.NewRat(19, 56)
	rp := big.NewRat(1111, 2222)
	rq := big.NewRat(1, 1)
	rq.Mul(rm, rn)
	fmt.Printf("Big Rat: %v\n", rq)
	rq.Mul(rm, rn).Add(rq, ro)
	fmt.Printf("Big Rat: %v\n", rq)
	rq.Mul(rm, rn).Add(rq, ro).Mul(rq, rp)
	fmt.Printf("Big Rat: %v\n", rq)
	var value []int
	value = []int{1,2,3}
	value = value[3:]
	if len(value) != 0 {
		fmt.Println(value)
	} else {
		fmt.Println("Ran out of stuff to do, exiting")
	}
}