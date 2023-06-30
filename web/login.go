package web

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	. "thinkPrinter/database"
	"thinkPrinter/entity"
	"thinkPrinter/tools"
	"time"
)

func Login(c *gin.Context) {
	// 数据库读取user
	var user entity.User
	// JSON解析到user
	var userDTO entity.User

	// 读取请求体中的数据
	err := c.BindJSON(&userDTO)
	if err != nil {
		log.Panicln(err)
	}

	// 加密密码
	userDTO.Password = tools.Encrypt(userDTO.Password)

	result := DB.Where("username = ? AND password = ?", userDTO.Username, userDTO.Password).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "用户名或密码错误",
		})
		return
	}

	// 生成token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["vip"] = user.Vip
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte("thinkPrinter"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "token生成错误",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"token":   tokenString,
	})
}
