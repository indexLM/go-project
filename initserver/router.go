package initserver

import (
	"github.com/gin-gonic/gin"
	"go-project/handler"
)

func routerCommonInit(Router *gin.RouterGroup) {
	AuthRouter := Router.Group("/auth")
	{
		AuthRouter.POST("/login", handler.Login)
	}
}
func routerSecurityInit(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("/auth")
	{
		BaseRouter.DELETE("/logout", handler.Logout)
	}
	OrderRouter := Router.Group("/order")
	{
		OrderRouter.GET("/list", handler.OrderList)
	}
}
