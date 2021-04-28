package handler

import (
	"github.com/gin-gonic/gin"
	"go-project/config"
	"go-project/model/req"
	"go-project/model/resp"
	"go-project/service"
)

func Login(c *gin.Context) {
	var req req.LoginRequest
	_ = c.ShouldBind(&req)
	if req.Username == "" || req.Password == "" {
		panic("参数错误")
	}
	passwordLogin, err := service.PasswordLogin(&req)
	if err != nil {
		resp.FailWithMessage(passwordLogin, c)
		return
	}
	resp.OkWithData(passwordLogin, c)
}
func Logout(c *gin.Context) {
	i := c.Keys["users"]
	o := i.(*config.CustomClaims)
	resp.OkWithData(o.UserId, c)
}
