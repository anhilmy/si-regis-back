package model

type Kategori struct {
	ID       int32  `json:"id"`
	Nama     string `json:"nama"`
	IsActive bool   `json:"is_active"`
}
