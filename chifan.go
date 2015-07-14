package main

import (
	"math/rand"
	// "os/exec"
	"time"
	// "os"
	 "fmt"
)

func main() {
	shop := []string{"黄焖鸡", "麻辣烫", "石锅拌饭", "拉面", "外卖"}
	fmt.Println(shop[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(shop))])
	//莫要瞎搞
	// cmd := exec.Command("cmd.exe", "/c", "shutdown -l")
	// cmd := exec.Command("cmd.exe", "/c", "pause")
	// cmd.Run()
	// cmd.Wait()
}
