package utils

import (
	"github.com/gin-gonic/gin"
	"go-project/config"
	"go-project/global"
	"time"
)

//生成令牌
func JwtGenerate(userId string, username string) (string, error) {
	claims := config.CustomClaims{
		UserId: userId,
		Name:   username,
	}
	// 签名生效时间
	claims.NotBefore = time.Now().Unix() - 1000
	// 过期时间 一小时
	claims.ExpiresAt = time.Now().Unix() + 3600
	//签发人
	claims.Issuer = "indexLm"
	jwt, err := global.MyJwt.Create(claims)
	if err != nil {
		return "生成令牌失败", err
	}
	return jwt, nil
}

func JwtVerification(tokenString string, c *gin.Context) error {
	parse, err := global.MyJwt.Parse(tokenString)
	if err != nil {
		return err
	}
	if c.Keys == nil {
		c.Keys = make(map[string]interface{}, 0)
	}
	c.Keys["users"] = parse
	return nil
}
