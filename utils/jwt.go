package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// 一些常量
var (
	TokenExpired     error  = errors.New("令牌已过期")
	TokenNotValidYet error  = errors.New("令牌尚未激活")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("无效的令牌")
	SignKey          string = "newtrekWang"
)

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	ID    string `json:"userId"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	jwt.StandardClaims
}

// 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// 获取signKey
func GetSignKey() string {
	return SignKey
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}
