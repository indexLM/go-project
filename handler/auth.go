package handler

import "github.com/gin-gonic/gin"

func RouterAuthInit(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("auth")
	{
		BaseRouter.DELETE("logout", nil)
		BaseRouter.POST("login", nil)
	}
}
