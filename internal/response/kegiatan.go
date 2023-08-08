package response

import "sireg/rest-api-kegiatan/internal/model"

type Kegiatan struct {
	ID       int32            `json:"id"`
	Judul    string           `json:"judul"`
	Desc     string           `json:"desc"`
	Kategori KegiatanKategori `json:"kategori"`
	NoSurat  string           `json:"no_surat"`
}

type KegiatanKategori struct {
	ID int32
}

func ConvertFromKegiatanModel(kategori *model.Kegiatan) Kegiatan {
	return Kegiatan{
		ID:      kategori.ID,
		Judul:   kategori.Judul.String,
		Desc:    kategori.Desc.String,
		NoSurat: kategori.NoSurat.String,
		Kategori: KegiatanKategori{
			ID: kategori.KategoriId.Int32,
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
