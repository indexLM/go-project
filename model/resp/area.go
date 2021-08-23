package resp

type DistrictListRes struct {
	CityId       *string `json:"cityId" db:"city_id"`
	CityName     *string `json:"CityName" db:"city_name"`
	DistrictId   *string `json:"districtId" db:"district_id"`
	DistrictName *string `json:"districtName" db:"district_name"`
}
