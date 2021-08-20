package req

import "time"

type OrderListReq struct {
	PatientName     string    `form:"patientName" json:"patientName"`
	PatientGender   uint64    `form:"patientGender" json:"PatientGender"`
	BranchId        uint64    `form:"branchId" json:"branchId"`
	AppointmentTime time.Time `form:"appointmentTime" json:"appointmentTime" time_format:"2006-01-02"`
	Status          uint64    `form:"status" json:"status"  `
}