package main

import (
	"fmt"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Say() {
	fmt.Println("你好啊")
	c.Ctx.WriteString("hello world")
}

func (c *MainController) Get() {
	fmt.Println("Test")
	c.Ctx.WriteString("Test")
}

func (c *MainController) Post() {
	fmt.Println("Post")
	c.Ctx.WriteString("Post")
}

func main() {
	web := beego.NewNamespace("/web",
		beego.NSRouter("/test", &MainController{}),
		beego.NSRouter("/*", &MainController{}, "get:Say"))

	beego.AddNamespace(web)
	beego.Run()
}
