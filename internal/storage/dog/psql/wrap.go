package psql

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type wrappedDB struct {
	DBTX
	b *Builder
}

func Wrap(db DBTX, b *Builder) DBTX {
	return &wrappedDB{
		DBTX: db,
		b:    b,
	}
}

func (w *wrappedDB) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	if w.b != nil {
		query, args = w.b.Build(query, args...)
	}

	return w.DBTX.Exec(ctx, query, args...)
}

func (w *wrappedDB) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	if w.b != nil {
		query, args = w.b.Build(query, args...)
	}

	return w.DBTX.Query(ctx, query, args...)
}

func (w *wrappedDB) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	if w.b != nil {
		query, args = w.b.Build(query, args...)
	}

	return w.DBTX.QueryRow(ctx, query, args...)
}
