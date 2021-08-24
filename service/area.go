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
func BranchList(cityId string, districtId string, branchName string) *[]resp.BranchListRes {
	sql := "select hb.id        as branch_id,\n       hb.name      as branch_name,\n       hb.address   as branch_address,\n       hb.longitude as lon,\n       hb.latitude  as lat\nfrom his_branch hb\n where city_id = ?\n  and district_id = ? \n  "
	res := make([]resp.BranchListRes, 0)
	if branchName != "" {
		sql = sql + " and hb.name like ?"
		err := global.MySqlx.Select(&res, sql, cityId, districtId, "%"+branchName+"%")
		if err != nil {
			panic(err.Error())
		}
	} else {
		err := global.MySqlx.Select(&res, sql, cityId, districtId)
		if err != nil {
			panic(err.Error())
		}
	}
	return &res
}
