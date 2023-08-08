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
