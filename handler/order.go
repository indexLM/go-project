package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-project/model/req"
	"go-project/model/resp"
	"go-project/service"
)

func OrderList(c *gin.Context) {
	r := new(req.OrderListReq)
	err := c.BindQuery(r)
	if err != nil {
		panic("参数错误")
	}
	branchId := c.Keys["branchId"]
	branchIdStr := fmt.Sprintf("%v", branchId)
	r.BranchId = branchIdStr
	res := service.GetOrderList(r)
	resp.OkWithData(res, c)
}
func OrderDetails(c *gin.Context) {
	orderNo := c.Query("orderNo")
	branchId := c.Keys["branchId"]
	branchIdStr := fmt.Sprintf("%v", branchId)
	res := service.GetOrderDetails(orderNo, branchIdStr)
	resp.OkWithData(res, c)
}
func ModifyOrderExt(c *gin.Context) {
	orderNo := c.PostForm("orderNo")
	ext := c.PostForm("ext")
	branchId := c.Keys["branchId"]
	branchIdStr := fmt.Sprintf("%v", branchId)
	service.ModifyOrderExt(orderNo, ext, branchIdStr)
	resp.Ok(c)
}
