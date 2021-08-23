package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-project/model/req"
	"go-project/model/resp"
	"go-project/service"
)

func PatientList(c *gin.Context) {
	r := new(req.PatientListReq)
	err := c.BindQuery(r)
	if err != nil {
		panic("参数错误")
	}
	branchId := c.Keys["branchId"]
	branchIdStr := fmt.Sprintf("%v", branchId)
	r.BranchId = branchIdStr
	res := service.GetPatientList(r)
	resp.OkWithData(res, c)
}
