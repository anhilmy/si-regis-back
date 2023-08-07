package model

type Kategori struct {
	ID       string `json:"id"`
	Nama     string `json:"nama"`
	IsActive bool   `json:"is_active"`
}
