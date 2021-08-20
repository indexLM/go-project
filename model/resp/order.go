package resp

type OrderListRes struct {
	PatientName   string `json:"patientName" db:"patient_name"`
	PatientGender uint64 `json:"PatientGender"  db:"patient_gender"`
	//BranchId        uint64    `json:"branchId" db:"branch_id"`
	//AppointmentTime time.Time `json:"appointmentTime" time_format:"2006-01-02"`
	//Status          uint64    `json:"status"  `
}
