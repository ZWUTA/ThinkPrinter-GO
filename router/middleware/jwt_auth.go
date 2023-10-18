package middleware

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	. "thinkprinter/models"
	"thinkprinter/tools"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if len(auth) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ERR.
				WithMsg("请登录").
				WithCode(http.StatusUnauthorized))
			return
		}

		if tokenType := auth[:6]; tokenType != "Bearer" {
			c.AbortWithStatusJSON(http.StatusBadRequest, ERR.
				WithMsg("token类型错误").
				WithCode(http.StatusBadRequest))
			return
		}
		token := auth[7:]
		slog.Debug("JWT Token", "token", token)

		claims, err := tools.ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ERR.
				WithMsg("token非法，请重新登录").
				WithCode(http.StatusUnauthorized))
			return
		}
		slog.Debug("JWT Claims", "claims", claims)
		username := claims.Audience
		c.Set("username", username[0])
		c.Next()
	}
}
