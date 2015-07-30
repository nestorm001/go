package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
	"io"
)

var num int16
var output string

func cf(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("method:", r.Method) //获取请求的方法else {

	shop := []string{"黄焖鸡", "麻辣烫", "石锅拌饭", "拉面", "外卖"}
	result := shop[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(shop))]
	if r.Method == "GET" {
		output = ""
	} else {
		num++
		output = "<pre>Hello there!\n今天吃" + result + "吧！\n</pre>"
		fmt.Printf("被访问了%d次\n", num)
		fmt.Println(result)
	}
	r.ParseForm() //解析参数，默认是不会解析的

	io.WriteString(w, "<html><head><meta http-equiv=\"Content-Type\"content=\"text/html; charset=utf-8\"/><title>今天吃什么？</title></head>" +
	"<body><form action=\"/\" method=\"post\"><input type=\"submit\" name=\"result\" value=\"随机一下\"></form>" + output + "</body></html>")
}

func main() {
	http.HandleFunc("/", cf) //设置访问的路由
	err := http.ListenAndServe(":8888", nil) //设置监听的端口
	num = 0
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
