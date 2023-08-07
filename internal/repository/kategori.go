package repository

import (
	"context"
	"database/sql"
	"sireg/rest-api-kegiatan/internal/model"
	"sireg/rest-api-kegiatan/pkg/dbcontext"
)

type KategoriRepo interface {
	GetAll(ctx context.Context) ([]model.Kategori, error)
}

type kategoriRepo struct {
	db *dbcontext.DB
}

func NewKategoriRepo(db *dbcontext.DB) KategoriRepo {
	return kategoriRepo{db}
}

func (r kategoriRepo) GetAll(ctx context.Context) ([]model.Kategori, error) {
	var kategori []model.Kategori
	err := r.db.With(ctx).Select().All(&kategori)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return kategori, err
}
