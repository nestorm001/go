package main

import (
    "time"
    "math/rand"
)

func main() {
	shop := []string{"黄焖鸡", "麻辣烫", "石锅拌饭", "拉面"}
	print(shop[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(shop))])    
}