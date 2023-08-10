package repository

import (
	"context"
	"database/sql"
	"sireg/rest-api-kegiatan/internal/model"
	"sireg/rest-api-kegiatan/pkg/dbcontext"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

type KategoriRepo interface {
	GetAll(ctx context.Context) ([]model.Kategori, error)
	Get(ctx context.Context, kategoriId int) (*model.Kategori, error)
	Create(ctx context.Context, kategori *model.Kategori) error
	Update(ctx context.Context, kategori *model.Kategori) error
	Delete(ctx context.Context, kategoriId int) error
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

func (r kategoriRepo) Get(ctx context.Context, kategoriId int) (*model.Kategori, error) {
	var kategori model.Kategori
	err := r.db.With(ctx).Select().Where(dbx.HashExp{"id": kategoriId}).One(&kategori)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &kategori, err
}
func (r kategoriRepo) Create(ctx context.Context, kategori *model.Kategori) error {
	err := r.db.DB().Model(kategori).Insert()
	return err
}
func (r kategoriRepo) Update(ctx context.Context, kategori *model.Kategori) error {
	err := r.db.DB().Model(kategori).Update()
	return err
}
func (r kategoriRepo) Delete(ctx context.Context, kategoriId int) error {
	err := r.db.DB().Model(model.Kategori{ID: int32(kategoriId)}).Delete()
	return err
}
