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
	AreaRouter := Router.Group("/area")
	{
		AreaRouter.GET("/district/list", handler.DistrictList)
		AreaRouter.GET("/branch/list", handler.BranchList)
		AreaRouter.POST("/branch/import", handler.BranchImport)
		AreaRouter.GET("/branch/adminReport", handler.BranchAdminReport)
		AreaRouter.POST("/doctor/import", handler.DoctorImport)
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
		OrderRouter.GET("/details", handler.OrderDetails)
		OrderRouter.POST("/modifyExt", handler.ModifyOrderExt)
	}
	PatientRouter := Router.Group("/patient")
	{
		PatientRouter.GET("/list", handler.PatientList)
	}
}
