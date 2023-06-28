package main

import (
	"fmt"
	"log"
	"net/http"
	"thinkPrinter/database"
	"thinkPrinter/tools"
	"thinkPrinter/web"
)

func main() {
	var port int = 8080
	var url = fmt.Sprintf("http://localhost:%d", port)
	// 检查sqlite数据库是否存在, 不存在则创建
	err := database.CheckSqlite()
	if err != nil {
		log.Fatal(err)
	}
	// 创建路由
	http.HandleFunc("/", web.Index)
	http.HandleFunc("/login", web.Login)
	http.HandleFunc("/signup", web.SignUp)
	log.Printf("Server is running at %s", url)
	tools.OpenExplorer(url)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Println("监听端口发生异常, 请检查端口是否被占用")
		panic(err)
	}
}
