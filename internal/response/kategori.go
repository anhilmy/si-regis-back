package response

import "sireg/rest-api-kegiatan/internal/model"

type Kategori struct {
	ID       int    `json:"id"`
	Nama     string `json:"nama"`
	IsActive bool   `json:"is_active"`
}

func ConvertKategori(kategori *model.Kategori) Kategori {
	return Kategori{
		ID:       kategori.ID,
		Nama:     kategori.Nama,
		IsActive: kategori.IsActive,
	}
}

func ConvertManyKategori(source []model.Kategori) []Kategori {
	var res []Kategori
	for _, kategori := range source {
		res = append(res, ConvertKategori(&kategori))
	}

	return res

}
