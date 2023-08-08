package request

type ReqKegiatan struct {
	Judul      string `json:"judul" binding:"required"`
	Desc       string `json:"desc"`
	KategoriId int    `json:"kategori_id" binding:"required"`
	NoSurat    string `json:"no_surat"`
}
