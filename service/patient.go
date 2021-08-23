package service

import (
	"go-project/global"
	"go-project/model/req"
	"go-project/model/resp"
	"go-project/utils"
	"strconv"
	"strings"
)

func GetPatientList(r *req.PatientListReq) *resp.PatientListRes {
	m := make(map[string]interface{})
	selectSql := "select distinct hp.patient_name," +
		"                hp.patient_gender," +
		"                hp.patient_age," +
		"                hot.patient_mobile," +
		"                hb.name         as branch_name," +
		"                hot.create_time as gmt_create "
	whereSql := "from his_order_treatment hot" +
		"         left join his_patient hp on hot.patient_id = hp.patient_id" +
		"         left join his_branch hb on hot.branch_id = hb.id " +
		"where 1 = 1 "
	name := r.PatientName
	patientName := strings.Trim(name, " ")
	if "" != patientName {
		whereSql = whereSql + " and hp.patient_name like :patient_name"
		m["patient_name"] = patientName + "%"
	}
	patientGender := r.PatientGender
	if 0 != patientGender {
		whereSql = whereSql + " and hp.patient_gender like :patient_gender"
		m["patient_gender"] = patientGender
	}
	patientMobile := r.PatientMobile
	if "" != patientMobile {
		whereSql = whereSql + " and hp.patient_mobile like :patient_mobile"
		m["patient_mobile"] = patientMobile
	}
	whereSql = whereSql + " and hb.id = :branch_id"
	m["branch_id"] = r.BranchId
	whereSql = whereSql + " and hot.goods_id = :goods_id"
	m["goods_id"] = 3
	currentPage := r.CurrentPage
	pageSize := r.PageSize
	start := (currentPage - 1) * pageSize
	limit := " limit " + strconv.FormatUint(start, 10) + " , " + strconv.FormatUint(pageSize, 10)
	utils.LogInfo("sql", selectSql+whereSql+limit)
	query, err := global.MySqlx.NamedQuery(selectSql+whereSql+limit, m)
	if err != nil {
		panic(err.Error())
	}
	patientListPoList := make([]resp.PatientListPo, 0)
	for query.Next() {
		r2 := new(resp.PatientListPo)
		err := query.StructScan(r2)
		if err != nil {
			panic(err.Error())
		}
		patientListPoList = append(patientListPoList, *r2)
	}
	countSql := "select count(1) "
	countRows, err := global.MySqlx.NamedQuery(countSql+whereSql, m)
	if err != nil {
		panic(err.Error())
	}
	res := new(resp.PatientListRes)
	for countRows.Next() {
		err := countRows.Scan(&res.Count)
		if err != nil {
			panic(err.Error())
		}
	}
	res.InfoList = patientListPoList
	return res
}
