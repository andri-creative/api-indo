package models

type District struct {
	ID        string `json:"id"`
	RegencyID string `json:"regency_id"`
	Name      string `json:"name"`
}
