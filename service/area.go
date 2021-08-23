package service

import (
	"go-project/global"
	"go-project/model/resp"
)

func DistrictList(cityName string) *[]resp.DistrictListRes {
	sql := "select distinct\n       hb.city_id,\n       hb.city_name,\n       hb.district_id,\n       hb.district_name\nfrom his_branch hb\nwhere city_name like ?"
	districtListRes := make([]resp.DistrictListRes, 0)
	err := global.MySqlx.Select(&districtListRes, sql, cityName+"%")
	if err != nil {
		panic(err.Error())
	}
	return &districtListRes
}
