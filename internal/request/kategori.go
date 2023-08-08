package request

import "sireg/rest-api-kegiatan/internal/model"

type ReqKategori struct {
	Nama     string `json:"nama" binding:"required"`
	IsActive bool   `json:"is_active" binding:"required"`
}

type PathKategoriID struct {
	KategoriId int `json:"id" uri:"kategoriId" binding:"required"`
}

func (k ReqKategori) ConvertToModel() *model.Kategori {
	return &model.Kategori{
		Nama:     k.Nama,
		IsActive: k.IsActive,
	}
}

func (k ReqKategori) ConvertToModelWithID(kategoriID int) *model.Kategori {
	return &model.Kategori{
		ID:       kategoriID,
		Nama:     k.Nama,
		IsActive: k.IsActive,
	}
}
