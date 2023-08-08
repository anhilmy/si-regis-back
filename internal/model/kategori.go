package model

type Kategori struct {
	ID       int    `json:"id"`
	Nama     string `json:"nama"`
	IsActive bool   `json:"is_active"`
}
