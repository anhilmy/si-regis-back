package request

import (
	"database/sql"
	"sireg/rest-api-kegiatan/internal/model"
	"strconv"
)

type ReqKegiatan struct {
	Judul    string              `json:"judul" binding:"required"`
	Desc     string              `json:"desc"`
	Kategori ReqKegiatanKategori `json:"kategori" binding:"required" db:"kategori_id"`
	NoSurat  string              `json:"no_surat"`
	Status   string              `json:"status" binding:"number,required,gte=0,lte,3"`
}

func (k ReqKegiatan) ConvertToModel() *model.Kegiatan {
	status, _ := strconv.ParseInt(k.Status, 0, 32)
	return &model.Kegiatan{
		Judul:      sql.NullString{String: k.Judul, Valid: true},
		Desc:       sql.NullString{String: k.Desc, Valid: true},
		KategoriId: sql.NullInt32{Int32: k.Kategori.ID, Valid: true},
		NoSurat:    sql.NullString{String: k.NoSurat, Valid: true},
		Status:     int32(status),
	}
}

func (k ReqKegiatan) ConvertToModelWithID(kegiatanId int32) *model.Kegiatan {
	status, _ := strconv.ParseInt(k.Status, 0, 32)
	return &model.Kegiatan{
		ID:         kegiatanId,
		Judul:      sql.NullString{String: k.Judul, Valid: true},
		Desc:       sql.NullString{String: k.Desc, Valid: true},
		KategoriId: sql.NullInt32{Int32: k.Kategori.ID, Valid: true},
		Kategori: model.Kategori{
			ID: k.Kategori.ID,
		},
		Status:  int32(status),
		NoSurat: sql.NullString{String: k.NoSurat, Valid: true},
	}
}

type ReqKegiatanKategori struct {
	ID int32 `json:"id" binding:"required"`
}

type PathKegiatanID struct {
	ID int32 `json:"id" binding:"required" uri:"id"`
}
