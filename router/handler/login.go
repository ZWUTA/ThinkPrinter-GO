package handler

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	. "thinkprinter/models"
	"thinkprinter/tools"
)

func Login(c *gin.Context) {
	var loginForm LoginForm
	var user User

	err := c.ShouldBindJSON(&loginForm)
	if err != nil {
		slog.Error("json绑定失败", "err", err)
		c.JSON(http.StatusBadRequest, ERR.WithMsg("请求体结构错误").
			WithCode(http.StatusBadRequest))
		return
	}

	loginForm.Password = tools.Encrypt(loginForm.Password)
	err = DB.Where("username = ? AND password = ?", loginForm.Username, loginForm.Password).
		Take(&user).Error
	if err != nil {
		slog.Error("查询错误", "err", err)
		c.JSON(http.StatusUnauthorized, ERR.WithMsg("用户名或密码不正确").
			WithCode(http.StatusUnauthorized))
		return
	}
	token, err := tools.CreateToken(user)
	if err != nil {
		slog.Error("生成token失败", "err", err)
		c.JSON(http.StatusInternalServerError, ERR)
		return
	}

	c.JSON(http.StatusOK, OK.WithData(gin.H{
		"token": token,
	}))
}
