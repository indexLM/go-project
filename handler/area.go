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
