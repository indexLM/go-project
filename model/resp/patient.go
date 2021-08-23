package resp

import "time"

type PatientListRes struct {
	Count    uint64          `json:"count"`
	InfoList []PatientListPo `json:"infoList"`
}
type PatientListPo struct {
	PatientName   *string   `json:"patientName" db:"patient_name"`
	PatientGender uint64    `json:"patientGender" db:"patient_gender"`
	PatientAge    uint64    `json:"patientAge" db:"patient_age"`
	PatientMobile *string   `json:"patientMobile" db:"patient_mobile"`
	BranchName    *string   `json:"branchName" db:"branch_name"`
	GmtCreate     time.Time `json:"gmtCreate" db:"gmt_create" time_format:"2006-01-02 15:04:05"`
}
