package model

import "database/sql"

type Kegiatan struct {
	ID       sql.NullInt32
	Judul    sql.NullString
	Desc     sql.NullString
	Kategori Kategori
	NoSurat  sql.NullString
}
