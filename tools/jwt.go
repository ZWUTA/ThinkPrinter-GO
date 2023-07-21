package tools

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"thinkPrinter/config"
	"thinkPrinter/entity"
	"time"
)

type MyClaims struct {
	Username string `json:"username"`
	Vip      bool   `json:"vip"`
	jwt.RegisteredClaims
}

// SignJWT 签发 JWT
func SignJWT(user entity.User) (string, error) {
	claims := MyClaims{
		Username: user.Username,
		Vip:      user.Vip,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.C.JWTExpiration) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.C.Security.JWTSecret)
	if err != nil {
		return "", fmt.Errorf("签署token失败: %v", err)
	}
	return tokenString, nil
}

// ParseJWT 解析 JWT
func ParseJWT(tokenString string) (*MyClaims, error) {
	var claims MyClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return config.C.Security.JWTSecret, nil
	})
	if err != nil {
		return nil, fmt.Errorf("解析token失败: %v", err)
	}
	if !token.Valid {
		return nil, fmt.Errorf("无效的token")
	}
	return &claims, nil
}
