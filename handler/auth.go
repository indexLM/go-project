package handler

import (
	"github.com/gin-gonic/gin"
	"go-project/config"
	"go-project/model/request"
	response "go-project/model/resopnse"
	"go-project/service"
)

func Login(c *gin.Context) {
	var req request.LoginRequest
	_ = c.ShouldBind(&req)
	if req.Username == "" || req.Password == "" {
		panic("参数错误")
	}
	passwordLogin, err := service.PasswordLogin(&req)
	if err != nil {
		response.FailWithMessage(passwordLogin, c)
		return
	}
	response.OkWithData(passwordLogin, c)
}
func Logout(c *gin.Context) {
	i := c.Keys["users"]
	o := i.(*config.CustomClaims)
	response.OkWithData(o.UserId, c)
}
