package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var num int16
var slice []string

func chifan(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("web_study\\generate.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm() //解析参数，默认是不会解析的
		num++
		fmt.Printf("被访问了%d次\n", num)
		shop := []string{"黄焖鸡", "麻辣烫", "石锅拌饭", "拉面", "外卖"}
		result := shop[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(shop))]
		//		fmt.Fprintf(w, "Hello there!\n") //这个写入到w的是输出到客户端的
		//		fmt.Fprintf(w, "今天吃" + result + "吧！\n")
		template.HTMLEscape(w, []byte("Hello there!\n今天吃"+result+"吧！\n"))
		//		t, _ := template.ParseFiles("web_study\\generate.html")
		//		t.Execute(w, nil)
		//		r.Form.Set("submit", result)
		fmt.Println(result)
		//		r.Form.Set("result", "Hello there!\n今天吃" + result + "吧！\n")
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("web_study\\login.html")
		t.Execute(w, token)
	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			//验证token的合法性
		} else {
			//不存在token报错
		}
		if len(r.Form["username"][0]) == 0 || len(r.Form["password"][0]) == 0 {
			fmt.Fprintf(w, "用户名或密码不能为空！")
		}
		fmt.Println("username length:", len(r.Form["username"][0]))
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端
	}
}

func main() {
	http.HandleFunc("/", chifan) //设置访问的路由
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":8888", nil) //设置监听的端口
	num = 0
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
