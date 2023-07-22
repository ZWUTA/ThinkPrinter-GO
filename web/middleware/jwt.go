package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"thinkPrinter/tools"
	"time"
)

// 白名单路径
var whiteList = []string{
	"/api/login",
	"/api/signup",
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 白名单
		for _, path := range whiteList {
			if path == c.Request.URL.Path {
				c.Next()
				return
			}
		}

		var m string

		token := c.GetHeader("Authorization")
		if token == "" {
			m = "请求未携带token，无权限访问"
		} else {
			claims, err := tools.ParseJWT(token)
			if err != nil {
				m = err.Error()
				log.Println(err)
			}

			if time.Now().Unix() > claims.ExpiresAt.Unix() {
				m = "token已过期"
			}
		}

		c.JSON(http.StatusUnauthorized, gin.H{
			"message": m,
		})
		c.Abort()
		return
	}
}
