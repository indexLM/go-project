package handler

import (
	"github.com/gin-gonic/gin"
	"go-project/model/resp"
	"go-project/service"
)

func DistrictList(c *gin.Context) {
	cityName := c.Query("cityName")
	if cityName == "" {
		panic("城市不能为空")
	}
	res := service.DistrictList(cityName)
	resp.OkWithData(res, c)
}
func BranchList(c *gin.Context) {
	cityId := c.Query("cityId")
	if cityId == "" {
		panic("城市不能为空")
	}
	districtId := c.Query("districtId")
	if districtId == "" {
		panic("区不能为空")
	}
	branchName := c.Query("branchName")
	res := service.BranchList(cityId, districtId, branchName)
	resp.OkWithData(res, c)
}
