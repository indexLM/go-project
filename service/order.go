package service

import (
	"go-project/global"
	"go-project/model/req"
	"go-project/model/resp"
	"go-project/utils"
	"strconv"
	"strings"
)

func GetOrderList(r *req.OrderListReq) *resp.OrderRes {
	m := make(map[string]interface{})
	selectList := "select hot.patient_name,hot.patient_gender "
	sql := "from his_order_treatment as hot inner join his_schedule_period_date_object as hspdo on hot.schedule_id=hspdo.id inner join his_schedule_date as hsd on hspdo.date_id=hsd.id where  1=1 "
	name := r.PatientName
	nameTrim := strings.Trim(name, " ")
	if "" != nameTrim {
		sql = sql + " and hot.patient_name like :patient_name"
		m["patient_name"] = "%" + nameTrim + "%"
	}
	gender := r.PatientGender
	if gender != 0 {
		sql = sql + " and hot.patient_gender = :patient_gender "
		m["patient_gender"] = gender
	}
	branchId := r.BranchId
	if branchId != 0 {
		sql = sql + " and hot.branch_id = :branch_id"
		m["branch_id"] = branchId
	}
	status := r.Status
	if status != 0 {
		sql = sql + " and hot.status = :status "
		m["status"] = status
	}
	appointmentTime := r.AppointmentTime
	if !appointmentTime.IsZero() {
		sql = sql + " and hsd.schedule_date >= :appointment_time "
		m["appointment_time"] = appointmentTime
	}
	currentPage := r.CurrentPage
	pageSize := r.PageSize
	start := (currentPage - 1) * pageSize
	limit := " limit " + strconv.FormatUint(start, 10) + " , " + strconv.FormatUint(pageSize, 10)
	utils.LogInfo("sql", sql)
	rows, err := global.MySqlx.NamedQuery(selectList+sql+limit, m)
	if err != nil {
		panic(err.Error())
	}
	orderListRes := make([]resp.OrderListRes, 0)
	for rows.Next() {
		r := new(resp.OrderListRes)
		err := rows.StructScan(r)
		if err != nil {
			panic(err.Error())
		}
		orderListRes = append(orderListRes, *r)
	}
	selectCount := "select count(1) "
	countQuery, err := global.MySqlx.NamedQuery(selectCount+sql, m)
	if err != nil {
		panic(err.Error())
	}
	res := new(resp.OrderRes)
	for countQuery.Next() {
		err := countQuery.Scan(&res.Count)
		if err != nil {
			panic(err.Error())
		}
	}
	res.InfoList = orderListRes
	return res
}
