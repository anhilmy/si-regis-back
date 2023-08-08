package kategori

import (
	"context"
	"sireg/rest-api-kegiatan/internal/model"
	"sireg/rest-api-kegiatan/internal/repository"
	"sireg/rest-api-kegiatan/internal/response"
)

type Service interface {
	GetAll(ctx context.Context) (*response.SuccessResponse, error)
}

type service struct {
	kategoriRepo repository.KategoriRepo
}

func NewService(kategoriRepo repository.KategoriRepo) Service {
	return service{kategoriRepo}
}

func (s service) GetAll(ctx context.Context) (*response.SuccessResponse, error) {
	kategori, err := s.kategoriRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var res = response.SuccessResponse{
		Message: "success",
		Data:    kategori,
	}
	return &res, nil
}

func (s service) Get(ctx context.Context, kategoriId int) (*response.SuccessResponse, error) {
	return nil, nil
}

func (s service) Create(ctx context.Context, kategori *model.Kategori) (*response.SuccessResponse, error) {
	return nil, nil
}

func (s service) Update(ctx context.Context, kategori *model.Kategori) (*response.SuccessResponse, error) {
	return nil, nil
}

func (s service) Delete(ctx context.Context, kategoriId int) (*response.SuccessResponse, error) {
	return nil, nil
}
