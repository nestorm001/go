package main

import (
	"fmt"
	"log"
	"net/http"
	"math/rand"
	"time"
)

func chifan(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Fprintf(w, "Hello there!\n") //这个写入到w的是输出到客户端的
	shop := []string{"黄焖鸡", "麻辣烫", "石锅拌饭", "拉面", "外卖"}
	fmt.Fprintf(w, "今天吃" + shop[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(shop))] + "吧！\n")
}

func main() {
	http.HandleFunc("/", chifan)       //设置访问的路由
	err := http.ListenAndServe(":8888", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
