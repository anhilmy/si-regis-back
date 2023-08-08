package request

import (
	"database/sql"
	"sireg/rest-api-kegiatan/internal/model"
)

type ReqKegiatan struct {
	Judul    string              `json:"judul" binding:"required"`
	Desc     string              `json:"desc"`
	Kategori ReqKegiatanKategori `json:"kategori" binding:"required" db:"kategori_id"`
	NoSurat  string              `json:"no_surat"`
}

func (k ReqKegiatan) ConvertToModel() *model.Kegiatan {
	return &model.Kegiatan{
		Judul:      sql.NullString{String: k.Judul, Valid: true},
		Desc:       sql.NullString{String: k.Desc, Valid: true},
		KategoriId: sql.NullInt32{Int32: k.Kategori.ID, Valid: true},
		Kategori: model.Kategori{
			ID: int(k.Kategori.ID),
		},
		NoSurat: sql.NullString{String: k.NoSurat, Valid: true},
	}
}

func (k ReqKegiatan) ConvertToModelWithID(kategoriID int32) *model.Kegiatan {
	return &model.Kegiatan{
		ID:         kategoriID,
		Judul:      sql.NullString{String: k.Judul, Valid: true},
		Desc:       sql.NullString{String: k.Desc, Valid: true},
		KategoriId: sql.NullInt32{Int32: k.Kategori.ID, Valid: true},
		Kategori: model.Kategori{
			ID: int(k.Kategori.ID),
		},
		NoSurat: sql.NullString{String: k.NoSurat, Valid: true},
	}
}

type ReqKegiatanKategori struct {
	ID int32 `json:"id" binding:"required"`
}

type PathKegiatanID struct {
	ID int32 `json:"id" binding:"required" uri:"id"`
}
