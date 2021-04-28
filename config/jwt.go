package config

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 一些常量
var (
	TokenExpired     error = errors.New("令牌已过期")
	TokenNotValidYet error = errors.New("令牌尚未激活")
	TokenMalformed   error = errors.New("That's not even a token")
	TokenInvalid     error = errors.New("无效的令牌")
)

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	UserId string `json:"userId"`
	Name   string `json:"name"`
	jwt.StandardClaims
}

// 生成token
func (j *Jwt) Create(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.SignKey))
}

// 解析token
func (j *Jwt) Parse(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.SignKey), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// 更新token
func (j *Jwt) Refresh(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.SignKey), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.Create(*claims)
	}
	return "", TokenInvalid
}
