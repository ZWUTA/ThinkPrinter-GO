package main

import (
	"embed"
	"fmt"
	"log"
	"thinkPrinter/config"
	"thinkPrinter/database"
	"thinkPrinter/tools"
	"thinkPrinter/web"
	"thinkPrinter/web/middleware"

	"github.com/gin-gonic/gin"
)

//go:embed static
var f embed.FS

func init() {
	// 初始化、迁移数据库
	err := database.InitDB()
	if err != nil {
		log.Panicln("数据库初始化失败", err)
	}
	// 传入静态文件
	web.F = f
}

func main() {
	url := fmt.Sprintf("%s:%d", config.C.Core.Bind, config.C.Core.Port)

	log.Printf("程序正在监听地址 %s", url)

	// 创建路由
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.HandleMethodNotAllowed = true

	// JWT 中间件
	r.Use(middleware.JWT())

	r.GET("/", web.Index)
	r.POST("/api/login", web.Login)
	r.POST("/api/signup", web.SignUp)

	// 打开浏览器
	tools.OpenBrowser(config.C.Core.Bind, config.C.Core.Port)

	// 监听地址
	err := r.Run(url)
	if err != nil {
		log.Println("监听端口发生异常, 请确保权限，并检查端口是否被占用", err)
	}
}
