package define

import "github.com/golang-jwt/jwt/v4"

type UserClaim struct {
	Id       uint
	Identity string
	Name     string
	jwt.RegisteredClaims
}

var JwtKey = "cloud-disk-key"
