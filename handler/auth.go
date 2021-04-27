package handler

import (
	"github.com/gin-gonic/gin"
	"go-project/model/request"
	response "go-project/model/resopnse"
	"go-project/service"
)

func RouterAuthInit(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("auth")
	{
		BaseRouter.DELETE("logout", nil)
		BaseRouter.POST("login", login)
	}
}

func login(c *gin.Context) {
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
