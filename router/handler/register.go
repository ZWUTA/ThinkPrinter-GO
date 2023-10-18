package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
	. "thinkprinter/models"
	"thinkprinter/tools"
)

func Register(c *gin.Context) {
	var loginForm LoginForm

	err := c.ShouldBindJSON(&loginForm)
	if err != nil {
		slog.Error("json绑定失败", "err", err)
		c.JSON(http.StatusBadRequest, ERR.WithMsg("请求体结构错误").
			WithCode(http.StatusBadRequest))
		return
	}

	loginForm.Password = tools.Encrypt(loginForm.Password)
	user := loginForm.ToUser()

	err = DB.Create(&user).Error
	if err != nil {
		slog.Error("新建用户失败", "err", err)
		// 当前上游有bug，无法判断约束冲突，TODO
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			c.JSON(http.StatusBadRequest, ERR.WithMsg("用户名已存在").
				WithCode(http.StatusBadRequest))
			return
		}
		c.JSON(http.StatusInternalServerError, ERR.WithMsg("注册失败"))
		return
	}

	c.JSON(http.StatusOK, OK)
}
