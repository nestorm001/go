package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
	"html/template"
)

var num int16

type Result struct {
	Output string
}

func cf(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println("method:", r.Method) //获取请求的方法else {
	t, _ := template.ParseFiles("web_study\\generate.html")  //解析模板文件

	if r.Method == "GET" {
		res := Result{Output:""}
		t.Execute(w, res)
	} else {
		num++
		fmt.Println(r.UserAgent())
		fmt.Printf("被访问了%d次\n", num)
		shop := []string{"黄焖鸡", "麻辣烫", "石锅拌饭", "拉面", "外卖"}
		randResult := shop[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(shop))]
		res := Result{Output:"Hello there!\n今天吃" + randResult + "吧！\n"}
		t.Execute(w, res)  //执行模板的merger操作
		fmt.Println(randResult)
	}
}

func main() {
	http.HandleFunc("/", cf) //设置访问的路由
	err := http.ListenAndServe(":8888", nil) //设置监听的端口
	num = 0
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
