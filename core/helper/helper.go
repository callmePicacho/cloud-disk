package helper

import (
	"cloud-disk/core/define"
	"context"
	"crypto/md5"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
	"math/rand"
	"net/http"
	"net/smtp"
	"net/url"
	"path"
	"time"
)

// Md5 加盐加密
func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s+define.PasswordSalt)))
}

// GenerateToken 生成 Token
func GenerateToken(id uint, identity, name string, second int) (string, error) {
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(second) * time.Second)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	signedString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return signedString, nil
}

// AnalyzeToken Token 解析
func AnalyzeToken(token string) (uc *define.UserClaim, err error) {
	uc = new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})
	if err != nil {
		return
	}
	if !claims.Valid {
		return uc, errors.New("token is invalid")
	}
	return uc, err
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

// CosUpload 上传文件到腾讯云
func CosUpload(r *http.Request) (string, error) {
	u, _ := url.Parse(define.TencentCosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		return "", err
	}
	key := GenerateUUID() + path.Ext(fileHeader.Filename)

	_, err = client.Object.Put(
		context.Background(), key, file, nil,
	)
	return define.TencentCosBucket + "/" + key, err
}
