package request

type ReqKategori struct {
	Nama     string `json:"nama" binding:"required"`
	IsActive bool   `json:"is_active" binding:"required"`
}
