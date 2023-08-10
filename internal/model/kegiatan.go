package model

import "database/sql"

type Kegiatan struct {
	ID         int32
	Judul      sql.NullString
	Desc       sql.NullString
	KategoriId sql.NullInt32 `db:"kategoriId"`
	Kategori   Kategori      `db:"-"`
	NoSurat    sql.NullString
}

type DTOSummaryKegiatan struct {
	Total        sql.NullInt32 `db:"total_kegiatan"`
	KategoriId   int32         `db:"kategoriId"`
	NamaKategori string        `db:"nama"`
	IsActive     string        `db:"is_active"`
}
