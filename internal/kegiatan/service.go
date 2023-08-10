package kegiatan

import (
	"context"
	"sireg/rest-api-kegiatan/internal/repository"
	"sireg/rest-api-kegiatan/internal/request"
	"sireg/rest-api-kegiatan/internal/response"
)

type Service interface {
	GetAll(ctx context.Context) ([]response.Kegiatan, error)
	Get(ctx context.Context, kegiatanId int32) (*response.Kegiatan, error)
	Create(ctx context.Context, kegiatan *request.ReqKegiatan) (*response.Kegiatan, error)
	Update(ctx context.Context, kegiatan *request.ReqKegiatan, kegiatanId int32) (*response.Kegiatan, error)
	Delete(ctx context.Context, kegiatanId int32) error
	Summary(ctx context.Context) ([]response.SummaryKegiatan, error)
}

type service struct {
	kegiatanRepo repository.KegiatanRepo
}

func NewService(kegiatanRepo repository.KegiatanRepo) Service {
	return service{kegiatanRepo}
}

func (s service) GetAll(ctx context.Context) ([]response.Kegiatan, error) {
	kegiatan, err := s.kegiatanRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var res = response.ConvertFromManyKegiatanModel(kegiatan)
	return res, nil
}

func (s service) Get(ctx context.Context, kegiatanId int32) (*response.Kegiatan, error) {
	kegiatan, err := s.kegiatanRepo.Get(ctx, kegiatanId)
	if err != nil {
		return nil, err
	}
	res := response.ConvertFromKegiatanModel(kegiatan)
	return &res, nil
}

func (s service) Create(ctx context.Context, request *request.ReqKegiatan) (*response.Kegiatan, error) {
	kegiatan := request.ConvertToModel()

	err := s.kegiatanRepo.Create(ctx, kegiatan)
	if err != nil {
		return nil, err
	}
	res := response.ConvertFromKegiatanModel(kegiatan)
	return &res, nil
}

func (s service) Update(ctx context.Context, request *request.ReqKegiatan, kegiatanId int32) (*response.Kegiatan, error) {
	kegiatan := request.ConvertToModelWithID(kegiatanId)

	err := s.kegiatanRepo.Update(ctx, kegiatan)
	if err != nil {
		return nil, err
	}
	res := response.ConvertFromKegiatanModel(kegiatan)
	return &res, nil
}

func (s service) Delete(ctx context.Context, kegiatanId int32) error {
	err := s.kegiatanRepo.Delete(ctx, kegiatanId)
	if err != nil {
		return err
	}
	return nil
}

func (s service) Summary(ctx context.Context) ([]response.SummaryKegiatan, error) {
	DTOSumm, err := s.kegiatanRepo.GetSummary(ctx)
	if err != nil {
		return nil, err
	}

	res := response.ConvertFromManyTOSummary(DTOSumm)
	return res, nil
}
