package helper

import (
	"cloud-disk/core/define"
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"net/smtp"
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

// MailSendCode 邮箱验证码发送
func MailSendCode(mail, code string) error {
	e := email.NewEmail()
	e.From = "cloud-disk<liyuanyue1996@163.com>"
	e.To = []string{"liyuanyue.cqucc@foxmail.com"}
	e.Subject = "验证码发送"
	e.HTML = []byte("您的验证码为：<h1>" + code + "</h1>")
	return e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "liyuanyue1996@163.com", define.MailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
}

// GenerateVerifyCode 生成验证码
func GenerateVerifyCode() (code string) {
	s := "1234567890"
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < define.CodeLength; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return
}

// GenerateUUID 生成UUID
func GenerateUUID() string {
	return uuid.NewV4().String()
}
