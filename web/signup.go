package web

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	. "thinkPrinter/database"
	"thinkPrinter/entity"
	"thinkPrinter/tools"
)

func SignUp(c *gin.Context) {
	// 数据库读取user
	var user entity.User
	// JSON解析到user
	var userDTO entity.User

	// 读取请求体中的数据
	err := c.BindJSON(&userDTO)
	if err != nil {
		log.Panicln(err)
	}

	// 查询用户是否存在
	result := DB.Where("username = ?", userDTO.Username).First(&user)
	if result.Error == nil {
		// 用户已存在
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "用户已存在",
		})

	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// 用户不存在，可以注册
		// 创建用户
		if userDTO.Username == "" || userDTO.Password == "" {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"message": "用户名或密码不能为空",
			})
			return
		}

		// 通过注册的用户默认不是VIP
		userDTO.Vip = false
		userDTO.Password = tools.Encrypt(userDTO.Password)

		result := DB.Create(&userDTO)
		if result.Error != nil {
			log.Panicln(result.Error)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "注册成功",
		})

	} else {
		// 其他错误
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "服务器错误",
		})
		log.Panicln(result.Error)

	}
}
