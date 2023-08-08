package kategori

import (
	"context"
	"sireg/rest-api-kegiatan/internal/repository"
	"sireg/rest-api-kegiatan/internal/request"
	"sireg/rest-api-kegiatan/internal/response"
)

type Service interface {
	GetAll(ctx context.Context) ([]response.Kategori, error)
	Get(ctx context.Context, kategoriId int) (*response.Kategori, error)
	Create(ctx context.Context, kategori *request.ReqKategori) (*response.Kategori, error)
	Update(ctx context.Context, kategori *request.ReqKategori, kategoriId int) (*response.Kategori, error)
	Delete(ctx context.Context, kategoriId int) error
}

type service struct {
	kategoriRepo repository.KategoriRepo
}

func NewService(kategoriRepo repository.KategoriRepo) Service {
	return service{kategoriRepo}
}

func (s service) GetAll(ctx context.Context) ([]response.Kategori, error) {
	kategori, err := s.kategoriRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var res = response.ConvertManyKategori(kategori)
	return res, nil
}

func (s service) Get(ctx context.Context, kategoriId int) (*response.Kategori, error) {
	kategori, err := s.kategoriRepo.Get(ctx, kategoriId)
	if err != nil {
		return nil, err
	}
	res := response.ConvertKategori(kategori)
	return &res, nil
}

func (s service) Create(ctx context.Context, request *request.ReqKategori) (*response.Kategori, error) {
	kategori := request.ConvertToModel()

	err := s.kategoriRepo.Create(ctx, kategori)
	if err != nil {
		return nil, err
	}
	res := response.ConvertKategori(kategori)
	return &res, nil
}

func (s service) Update(ctx context.Context, request *request.ReqKategori, kategoriId int) (*response.Kategori, error) {
	kategori := request.ConvertToModelWithID(kategoriId)

	err := s.kategoriRepo.Update(ctx, kategori)
	if err != nil {
		return nil, err
	}
	res := response.ConvertKategori(kategori)
	return &res, nil
}

func (s service) Delete(ctx context.Context, kategoriId int) error {
	err := s.kategoriRepo.Delete(ctx, kategoriId)
	if err != nil {
		return err
	}
	return nil
}
