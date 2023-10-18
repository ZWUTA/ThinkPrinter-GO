package router

import (
	"github.com/gin-gonic/gin"
	"thinkprinter/router/handler"
	"thinkprinter/router/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.HandleMethodNotAllowed = true

	// 分组注册路由
	root := r.Group("/")
	{
		root.GET("/", handler.Index)
	}
	api := r.Group("/api")
	{
		api.POST("/register", handler.Register)
		api.POST("/login", handler.Login)
	}
	protected := r.Group("/api/protected")
	{
		protected.Use(middleware.JWTAuth())

		protected.GET("/ping", handler.Ping)
		protected.POST("/upload", handler.Upload)
	}

	return r
}
