package req

type PatientListReq struct {
	PatientName   string `form:"patientName" json:"patientName"`
	PatientGender uint64 `form:"patientGender" json:"patientGender"`
	PatientMobile string `form:"patientMobile" json:"patientMobile"`
	BranchId      string `form:"branchId" json:"branchId"`
	PageSize      uint64 `form:"pageSize" json:"pageSize"`
	CurrentPage   uint64 `form:"currentPage" json:"currentPage"`
}
