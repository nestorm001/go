package main

import (
	. "fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		Println(s)
	}
}

func main() {
	Print("乱码用fmt就好了\n")
	go say("world") //开一个新的Goroutines执行
	say("hello") //当前Goroutines执行
	say("!")
	Println(runtime.NumCPU())
}