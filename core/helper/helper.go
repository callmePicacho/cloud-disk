package helper

import (
	"cloud-disk/core/define"
	"crypto/md5"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	Salt = "lyyyyy"
)

// Md5 加盐加密
func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s+Salt)))
}

// GenerateToken 生成 Token
func GenerateToken(id uint, identity, name string) (string, error) {
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)), // 默认七天过期时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	signedString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return signedString, nil
}
