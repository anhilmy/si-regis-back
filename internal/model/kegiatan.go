package model

import "database/sql"

type Kegiatan struct {
	ID         int32
	Judul      sql.NullString
	Desc       sql.NullString
	KategoriId sql.NullInt32 `db:"kategoriId"`
	Kategori   Kategori
	NoSurat    sql.NullString
	Status     int32
}

type DTOSummaryKegiatan struct {
	Total        sql.NullInt32 `db:"total_kegiatan"`
	KategoriId   int32         `db:"kategoriId"`
	NamaKategori string        `db:"nama"`
	IsActive     string        `db:"is_active"`
}

type StatusKegiatan int

const (
	PRAPELAKSANAAN StatusKegiatan = iota
	FIELDWORK
	PEMBUATAN_LAPORAN
	RILIS
)

var ListStatusKegiatan = map[int32]string{
	0: "PRAPELAKSANAAN",
	1: "FIELDWORK",
	2: "PEMBUATAN_LAPORAN",
	3: "RILIS",
}
