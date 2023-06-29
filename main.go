package main

import (
	"fmt"
	"log"
	"net/http"
	"thinkPrinter/database"
	"thinkPrinter/tools"
	"thinkPrinter/web"
)

var (
	port int
	url  string
)

func init() {
	port = 8080
	url = fmt.Sprintf("localhost:%d", port)

	// 检查sqlite数据库是否存在, 不存在则创建
	err := database.CheckSqlite()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// 创建路由
	http.HandleFunc("/", web.Index)
	http.HandleFunc("/login", web.Login)
	http.HandleFunc("/signup", web.SignUp)
	log.Printf("Server is running at %s", url)
	// 打开浏览器
	tools.OpenExplorer(url)
	// 监听端口
	err := http.ListenAndServe(url, nil)
	if err != nil {
		log.Println("监听端口发生异常, 请确保权限，并检查端口是否被占用")
		log.Panic(err)
	}
}
