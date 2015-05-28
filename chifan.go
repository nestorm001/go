package main

import (
    "time"
    "math/rand"
)

func main() {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
	shop := []string{"黄焖鸡", "麻辣烫", "石锅拌饭", "拉面"}
	print(shop[r.Intn(len(shop))])    
}