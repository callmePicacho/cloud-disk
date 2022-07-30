package test

import (
	"cloud-disk/core/define"
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func TestEmail(t *testing.T) {
	e := email.NewEmail()
	e.From = "Get <liyuanyue1996@163.com>"
	e.To = []string{"liyuanyue.cqucc@foxmail.com"}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("你的验证码为：<h1>123456</h1>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "liyuanyue1996@163.com", define.MailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		t.Fatal(err)
	}
}
