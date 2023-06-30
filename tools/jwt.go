package tools

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

// SignJWT 签发 JWT
func SignJWT(secretKey string, claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("签署token失败: %v", err)
	}
	return tokenString, nil
}

// VerifyJWT 验证 JWT
func VerifyJWT(tokenString string, secretKey string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("解析token失败: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("token非法")
	}

	return claims, nil
}
