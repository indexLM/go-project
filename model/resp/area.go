package resp

type DistrictListRes struct {
	CityId       *string `json:"cityId" db:"city_id"`
	CityName     *string `json:"cityName" db:"city_name"`
	DistrictId   *string `json:"districtId" db:"district_id"`
	DistrictName *string `json:"districtName" db:"district_name"`
}
type BranchListRes struct {
	BranchId      uint64  `json:"branchId" db:"branch_id"`
	BranchName    *string `json:"branchName" db:"branch_name"`
	BranchAddress *string `json:"branchAddress" db:"branch_address"`
	Lon           *string `json:"lon" db:"lon"`
	Lat           *string `json:"lat" db:"lat"`
}
