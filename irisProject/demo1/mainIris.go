package main

import (
	"fmt"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	// 这里的要么是绝对路径，要么就在main.go文件所在目录执行程序
	htmlEngine := iris.HTML("./templates", ".html")
	app.RegisterView(htmlEngine)
	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Hello World --from iris")
	})

	app.Get("/hello", func(ctx iris.Context) {
		ctx.ViewData("title", "测试页面")
		ctx.ViewData("content", "Hello World")
		err:=ctx.View("hello.html")
		if err!=nil{
			fmt.Println(err.Error())
		}
	})

	app.Run(iris.Addr(":9090"))
}
