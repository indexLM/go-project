package service

import (
	"go-project/global"
	"go-project/model/req"
	"go-project/model/resp"
	"go-project/utils"
)

func GetOrderList(r *req.OrderListReq) *[]resp.OrderListRes {
	name := r.PatientName
	m := make(map[string]interface{})
	m["patient_name"] = name
	where := utils.SqlFormatWhere(m)
	var sql = "select patient_name,patient_gender from his_order_treatment " + where
	rows, err := global.MySqlx.NamedQuery(sql, m)

	if err != nil {
		panic(err.Error())
	}
	res := make([]resp.OrderListRes, 0)
	for rows.Next() {
		r := new(resp.OrderListRes)
		err := rows.StructScan(r)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, *r)
	}
	return &res
}
