package main

import (
    "time"
    "math/rand"
	"os/exec"
	// "os"
	// "fmt"
)

func main() {
	shop := []string{"黄焖鸡", "麻辣烫", "石锅拌饭", "拉面", "外卖"}
	print(shop[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(shop))])
	//莫要瞎搞
    // cmd := exec.Command("cmd.exe", "/c", "shutdown -l")
    // cmd.Run()
}