package handler

import (
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
	res := service.GetOrderList(r)
	resp.OkWithData(res, c)
}
