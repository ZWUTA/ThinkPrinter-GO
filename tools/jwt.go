package tools

import (
	"errors"
	"strconv"
	"thinkprinter/models"
	"time"

	. "github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(models.C.Security.JWTSecret)

func CreateToken(user models.User) (string, error) {
	claims := RegisteredClaims{
		Issuer:   "ThinkPrinter-GO",
		Subject:  "Login",
		Audience: ClaimStrings{user.Username},
		ExpiresAt: NewNumericDate(time.Now().
			Add(time.Duration(models.C.Security.JWTExpiration) * time.Second)),
		NotBefore: NewNumericDate(time.Now()),
		IssuedAt:  NewNumericDate(time.Now()),
		ID:        strconv.Itoa(int(user.ID)),
	}

	token := NewWithClaims(SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ParseToken(tokenString string) (*RegisteredClaims, error) {
	token, err := ParseWithClaims(tokenString, &RegisteredClaims{}, func(token *Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*RegisteredClaims)
	if !ok || !token.Valid {
		return nil, errors.New("token不合法")
	}

	return claims, nil
}
