package repository

import (
	"context"
	"database/sql"
	constant "sireg/rest-api-kegiatan/internal/const"
	"sireg/rest-api-kegiatan/internal/model"
	"sireg/rest-api-kegiatan/pkg/dbcontext"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

type KegiatanRepo interface {
	GetAll(ctx context.Context) ([]model.Kegiatan, error)
	Get(ctx context.Context, kegiatanId int32) (*model.Kegiatan, error)
	Create(ctx context.Context, kegiatan *model.Kegiatan) error
	Update(ctx context.Context, kegiatan *model.Kegiatan) error
	Delete(ctx context.Context, kegiatanId int32) error
	GetSummary(ctx context.Context) ([]model.DTOSummaryKegiatan, error)
}

type kegiatanRepo struct {
	db *dbcontext.DB
}

func NewKegiatanRepo(db *dbcontext.DB) KegiatanRepo {
	return kegiatanRepo{db}
}

func (r kegiatanRepo) GetAll(ctx context.Context) ([]model.Kegiatan, error) {
	var kegiatan []model.Kegiatan
	err := r.db.With(ctx).Select("kegiatan.*", "kategori.id as kategori.id", "kategori.nama as kategori.nama").From(constant.TableKegiatan).LeftJoin(constant.TableKategori, dbx.NewExp(`kategori.id = kegiatan."kategoriId"`)).All(&kegiatan)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return kegiatan, err
}

func (r kegiatanRepo) Get(ctx context.Context, kegiatanId int32) (*model.Kegiatan, error) {
	var kegiatan model.Kegiatan
	err := r.db.With(ctx).Select().Where(dbx.HashExp{"id": sql.NullInt32{Int32: kegiatanId, Valid: true}}).One(&kegiatan)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &kegiatan, err
}
func (r kegiatanRepo) Create(ctx context.Context, kegiatan *model.Kegiatan) error {
	err := r.db.DB().Model(kegiatan).Insert()
	return err
}
func (r kegiatanRepo) Update(ctx context.Context, kegiatan *model.Kegiatan) error {
	err := r.db.DB().Model(kegiatan).Update()
	return err
}
func (r kegiatanRepo) Delete(ctx context.Context, kategoriId int32) error {
	err := r.db.DB().Model(&model.Kegiatan{ID: kategoriId}).Delete()
	return err
}

func (r kegiatanRepo) GetSummary(ctx context.Context) ([]model.DTOSummaryKegiatan, error) {
	var res []model.DTOSummaryKegiatan
	err := r.db.DB().Select("COUNT(*) as total_kegiatan", "kategori.id as kategoriId", "kategori.nama", "is_active").From("kategori").LeftJoin("kegiatan", dbx.NewExp(`kategori.id = kegiatan."kategoriId"`)).GroupBy("kategori.id", "kategori.nama", "is_active").OrderBy("kategori.id").All(&res)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return res, err
}
