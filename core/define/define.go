package define

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
)

type UserClaim struct {
	Id       uint
	Identity string
	Name     string
	jwt.RegisteredClaims
}

var JwtKey = "cloud-disk-key"

// MailPassword 从环境变量中读取的 email 密码
var MailPassword = os.Getenv("MailPassword")

// CodeLength 验证码长度
var CodeLength = 6

// CodeExpire 验证码过期时间，单位 s
var CodeExpire = 300

// TencentSecretKey 腾讯云对象存储
var TencentSecretKey = os.Getenv("TencentSecretKey")
var TencentSecretID = os.Getenv("TencentSecretID")

// TencentCosBucket 桶名称
var TencentCosBucket = "https://cloud-disk-1259459460.cos.ap-guangzhou.myqcloud.com"

// PageSize 分页的默认参数
var PageSize = 20
