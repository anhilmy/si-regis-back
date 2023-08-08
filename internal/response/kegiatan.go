package response

import "sireg/rest-api-kegiatan/internal/model"

type Kegiatan struct {
	ID       int32
	Judul    string
	Desc     string
	Kategori KegiatanKategori `json:"kategori"`
	NoSurat  string
}

type KegiatanKategori struct {
	ID int32
}

func ConvertFromKegiatanModel(kategori *model.Kegiatan) Kegiatan {
	return Kegiatan{
		ID:      kategori.ID.Int32,
		Judul:   kategori.Judul.String,
		Desc:    kategori.Desc.String,
		NoSurat: kategori.NoSurat.String,
		Kategori: KegiatanKategori{
			ID: int32(kategori.Kategori.ID),
		},
	}
}

func ConvertFromManyKegiatanModel(source []model.Kegiatan) []Kegiatan {
	var res []Kegiatan
	for _, kegiatan := range source {
		res = append(res, ConvertFromKegiatanModel(&kegiatan))
	}

	return res

}
