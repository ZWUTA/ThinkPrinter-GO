package main

import (
	"fmt"
	"log"
	"net/http"
	"thinkPrinter/database"
	"thinkPrinter/tools"
	"thinkPrinter/web"
)

const (
	// 绑定端口
	port = 8080
	// 监听地址
	bind = "0.0.0.0"
)

func init() {
	// 初始化、迁移数据库
	err := database.InitDB()
	if err != nil {
		log.Panicln("数据库初始化失败", err)
	}
}

func main() {
	url := fmt.Sprintf("%s:%d", bind, port)

	// 创建路由
	http.HandleFunc("/", web.Index)
	http.HandleFunc("/login", web.Login)
	http.HandleFunc("/signup", web.SignUp)
	log.Printf("Server is running at %s", url)
	// 打开浏览器
	tools.OpenBrowser(bind, port)
	// 监听端口
	err := http.ListenAndServe(url, nil)
	if err != nil {
		log.Println("监听端口发生异常, 请确保权限，并检查端口是否被占用")
		log.Panic(err)
	}
}
