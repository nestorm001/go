package main

import "time"
import "math"
import "fmt"

func main() {
	num := 65535.0
	t1 := time.Now()
	fmt.Printf("the sqrt of %f is %f", num, getsqrt(num))
	t2 := time.Now()
	fmt.Printf("\nexecuted in %d", t2.Sub(t1))
	fmt.Printf("\nthe sqrt of %f is %f", num, math.Sqrt(num))
	t3 := time.Now()
	fmt.Printf("\nexecuted in %d", t3.Sub(t2))
}

func getsqrt(x float64) float64 {
	z := x

	temp := x / 2
	for math.Abs(temp - z) > 0.000000001 {
		temp = z
		z = z - (z * z - x) / (2 * x)
	}
	return z
}