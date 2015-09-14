package main
import "fmt"

func main() {
	mf := map[int]func() int{
		1: func() int { return getSum(1) },
		2: func() int { return getSum(2) },
		5: func() int { return getSum(5) },
	}
	fmt.Println(mf)
	fmt.Println(mf[5])
	fmt.Println(mf[4])
	fmt.Println(mf[3])
	fmt.Println(mf[2])
	fmt.Println(mf[1])
}

func getSum(num int) int {
	sum := 0;
	for i := 0; i < num; i++ {
		sum += i
	}
	return sum
}