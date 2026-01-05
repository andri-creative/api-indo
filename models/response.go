package models

type RegencyByProvinceResponse struct {
	Province  Province  `json:"province"`
	Regencies []Regency `json:"regencies"`
}
