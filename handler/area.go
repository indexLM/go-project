package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-project/model/resp"
	"go-project/service"
	"net/http"
	"net/url"
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

func BranchImport(c *gin.Context) {
	file, err := c.FormFile("excel")
	if err != nil {
		panic(err.Error())
	}
	err = service.BranchImport(file)
	if err != nil {
		panic(err.Error())
	}
	resp.Ok(c)
}
func BranchAdminReport(c *gin.Context) {
	name := c.Query("name")
	data := service.BranchAdminReport(name)
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", url.QueryEscape("账号.xlsx")))
	c.Writer.Header().Set("Content-Type", "application/octet-stream")
	c.Data(http.StatusOK, "application/octet-stream", data)
}
func DoctorImport(c *gin.Context) {
	file, err := c.FormFile("excel")
	if err != nil {
		panic(err.Error())
	}
	err = service.DoctorImport(file)
	if err != nil {
		panic(err.Error())
	}
	resp.Ok(c)
}
