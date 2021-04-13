package handler

import (
	"github.com/gin-gonic/gin"
	"go-project/model/request"
	response "go-project/model/resopnse"
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
		response.FailWithMessage("参数错误", c)
		return
	}
	response.OkWithData(req.Username+"你好,你的密码是："+req.Password, c)
}
