package main
import (
	"fmt"
	"time"
)

var zero_num = 0
var one_num = 0
var two_num = 0
const total_num = 100000000

func main() {
	t1 := time.Now()
	ch := make(chan int, 1000)
	j := 0
	for {
		select {
		case ch <- 0:
		case ch <- 1:
		case ch <- 2:
		}
		i := <-ch
		switch i {
		case 0:
			zero_num++
			break
		case 1:
			one_num++
			break
		case 2:
			two_num++
			break
		}
		if j > total_num {
			break;
		}
		j++;
	}
	fmt.Println("Zero: ", zero_num)
	fmt.Println("One : ", one_num)
	fmt.Println("Two : ", two_num)
	fmt.Println("cost time:", time.Now().Sub(t1))

}