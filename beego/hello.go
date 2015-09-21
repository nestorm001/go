package main

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.TplNames = "./generate.html"
}

func main() {
	beego.Router("/", &MainController{})
	beego.Run()
}