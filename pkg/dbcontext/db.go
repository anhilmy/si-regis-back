package dbcontext

import (
	"context"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

type DB struct {
	db *dbx.DB
}

func New(db *dbx.DB) *DB {
	return &DB{db}
}

func (db *DB) DB() *dbx.DB {
	return db.db
}

func (db *DB) With(ctx context.Context) dbx.Builder {
	return db.db.WithContext(ctx)
}
