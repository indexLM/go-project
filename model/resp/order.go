package resp

import "time"

type OrderListRes struct {
	OrderNo       *string `json:"orderNo" db:"order_no"`
	PatientName   *string `json:"patientName" db:"patient_name"`
	PatientGender uint64  `json:"PatientGender"  db:"patient_gender"`
	BranchName    *string `json:"branchName" db:"branch_name"`
	BookingTime   *string `json:"bookingTime" db:"booking_time"`
	Status        uint64  `json:"status" db:"status" `
}
type OrderRes struct {
	Count    uint64         `json:"count"`
	InfoList []OrderListRes `json:"infoList"`
}
type OrderDetailsRes struct {
	OrderNo       *string    `json:"orderNo" db:"order_no"`
	PatientName   *string    `json:"patientName" db:"patient_name"`
	Status        uint64     `json:"status" db:"status" `
	Described     *string    `json:"described" db:"described"`
	BranchName    *string    `json:"branchName" db:"branch_name"`
	BookingTime   *string    `json:"bookingTime" db:"booking_time"`
	GmtCreate     *time.Time `json:"gmtCreate" db:"gmt_create" time_format:"2006-01-02 15:04:05"`
	PatientMobile *string    `json:"patientMobile" db:"patient_mobile"`
	Ext           *string    `json:"ext" db:"ext"`
}
