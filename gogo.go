package main
import (
	"fmt"
	"time"
)


func main() {
	timeFormat := "20060102"
	fileDate, _ := time.Parse(timeFormat, "20160222")
	todayDate, _ := time.Parse(timeFormat, "20160229")
	fmt.Println(todayDate.Sub(fileDate).Hours())

	go print1()
	go print2()
	time.Sleep(5 * time.Hour)
}

func print1() {
	for {
		fmt.Println("1")
		time.Sleep(1 * time.Second)
	}
}

func print2() {
	for {
		fmt.Println("2")
		time.Sleep(2 * time.Second)
	}
}

