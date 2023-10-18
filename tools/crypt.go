package tools

import (
	"crypto/sha256"
	"encoding/hex"
	"thinkprinter/models"
)

func Encrypt(password string) string {
	salt := models.C.Security.PasswordSalt
	// 创建 SHA256 哈希对象
	hash := sha256.New()

	// 将密码转换为字节数组并写入哈希对象
	hash.Write([]byte(password + salt))

	// 计算哈希值并获取结果
	hashBytes := hash.Sum(nil)

	// 将哈希值转换为十六进制字符串
	hashString := hex.EncodeToString(hashBytes)

	return hashString
}
