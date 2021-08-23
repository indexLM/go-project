package service

import (
	"go-project/global"
	"go-project/model/req"
	"go-project/model/resp"
	"go-project/utils"
	"strconv"
	"strings"
)

// GetOrderList 查询订单列表
func GetOrderList(r *req.OrderListReq) *resp.OrderRes {
	m := make(map[string]interface{})
	selectList := "select hot.order_no, hot.patient_name,\n" +
		"       hot.patient_gender,\n" +
		" 		hb.name as branch_name,\n" +
		"       hot.booking_time,\n" +
		"       hot.status "
	sql := "from his_order_treatment as hot " +
		"inner join his_schedule_period_date_object as hspdo on hot.schedule_id=hspdo.id " +
		"inner join his_schedule_date as hsd on hspdo.date_id=hsd.id " +
		"left join his_branch hb on hb.id=hot.branch_id " +
		"where  goods_id = 3 "
	name := r.PatientName
	nameTrim := strings.Trim(name, " ")
	if "" != nameTrim {
		sql = sql + " and hot.patient_name like :patient_name"
		m["patient_name"] = nameTrim + "%"
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
	utils.LogInfo("sql", selectList+sql+limit)
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

// GetOrderDetails 查询订单详情
func GetOrderDetails(orderNo string, branchId string) *resp.OrderDetailsRes {
	sql := "select hot.order_no,       hot.patient_name,       hot.status,       hotc.described,       hb.name         as branch_name,       hot.booking_time,       hot.create_time as gmt_create,       hot.patient_mobile,       hot.ext from his_order_treatment as hot         inner join his_order_treatment_case hotc on hot.order_no = hotc.order_no         inner join his_branch hb on hot.branch_id = hb.id where hot.order_no = ?   and hot.branch_id = ?"
	res := new(resp.OrderDetailsRes)
	err := global.MySqlx.Get(res, sql, orderNo, branchId)
	if err != nil {
		panic(err.Error())
	}
	return res
}

// ModifyOrderExt 修改订单备注
func ModifyOrderExt(orderNo string, ext string, branchId string) {
	if orderNo == "" {
		panic("订单号不能为空")
	}
	sql := "update his_order_treatment set ext=? where order_no=? and branch_id=? limit 1;"
	result, err := global.MySqlx.Exec(sql, ext, orderNo, branchId)
	if err != nil {
		panic(err.Error())
	}
	_, err = result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
}
