package response

import "sireg/rest-api-kegiatan/internal/model"

type Kegiatan struct {
	ID       int32            `json:"id"`
	Judul    string           `json:"judul"`
	Desc     string           `json:"desc"`
	Kategori KegiatanKategori `json:"kategori"`
	NoSurat  string           `json:"no_surat"`
	Status   string           `json:"status"`
}

type KegiatanKategori struct {
	ID   int32  `json:"ID"`
	Nama string `json:"nama"`
}

func ConvertFromKegiatanModel(kegiatan *model.Kegiatan) Kegiatan {
	return Kegiatan{
		ID:      kegiatan.ID,
		Judul:   kegiatan.Judul.String,
		Desc:    kegiatan.Desc.String,
		NoSurat: kegiatan.NoSurat.String,
		Kategori: KegiatanKategori{
			ID:   kegiatan.Kategori.ID,
			Nama: kegiatan.Kategori.Nama,
		},
		Status: model.ListStatusKegiatan[kegiatan.Status],
	}
}

func ConvertFromManyKegiatanModel(source []model.Kegiatan) []Kegiatan {
	var res []Kegiatan
	for _, kegiatan := range source {
		res = append(res, ConvertFromKegiatanModel(&kegiatan))
	}

	return res

}

type SummaryKegiatan struct {
	TotalKegiatan int32  `db:"total_kegiatan" json:"total_kegiatan"`
	KategoriId    int32  `db:"kategori.id" json:"id"`
	NamaKategori  string `db:"nama" json:"nama"`
	IsActive      string `db:"is_active" json:"is_active"`
}

func ConvertFromDTOSummary(DTOSumm model.DTOSummaryKegiatan) *SummaryKegiatan {
	return &SummaryKegiatan{
		TotalKegiatan: DTOSumm.Total.Int32,
		KategoriId:    DTOSumm.KategoriId,
		NamaKategori:  DTOSumm.NamaKategori,
		IsActive:      DTOSumm.IsActive,
	}
}

func ConvertFromManyTOSummary(DTOSumms []model.DTOSummaryKegiatan) []SummaryKegiatan {
	var res []SummaryKegiatan
	for _, summ := range DTOSumms {
		res = append(res, *ConvertFromDTOSummary(summ))
	}
	return res
}
