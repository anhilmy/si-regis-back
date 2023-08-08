package request

import "sireg/rest-api-kegiatan/internal/model"

type ReqKategori struct {
	Nama     string `json:"nama" binding:"required"`
	IsActive string `json:"is_active" binding:"required,boolean" example:"false"`
}

type PathKategoriID struct {
	KategoriId int `json:"id" uri:"kategoriId" binding:"required"`
}

func (k ReqKategori) ConvertToModel() *model.Kategori {
	var isActive bool
	if k.IsActive == "true" {
		isActive = true
	} else {
		isActive = false
	}
	return &model.Kategori{
		Nama:     k.Nama,
		IsActive: isActive,
	}
}

func (k ReqKategori) ConvertToModelWithID(kategoriID int) *model.Kategori {
	var isActive bool
	if k.IsActive == "true" {
		isActive = true
	} else {
		isActive = false
	}
	return &model.Kategori{
		ID:       kategoriID,
		Nama:     k.Nama,
		IsActive: isActive,
	}
}
