package sqlc

import (
	"context"
	"database/sql"
)

type (
	wrappedDBTX struct {
		db DBTX
	}
)

func Wrap(db DBTX) DBTX {
	return &wrappedDBTX{
		db: db,
	}
}

func (w *wrappedDBTX) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	if b, isOk := GetBuilder(ctx); isOk {
		query, args = b.Build(query, args...)
	}
	return w.db.ExecContext(ctx, query, args...)
}

func (w *wrappedDBTX) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return w.db.PrepareContext(ctx, query)
}

func (w *wrappedDBTX) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	if b, isOk := GetBuilder(ctx); isOk {
		query, args = b.Build(query, args...)
	}
	return w.db.QueryContext(ctx, query, args...)
}

func (w *wrappedDBTX) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	if b, isOk := GetBuilder(ctx); isOk {
		query, args = b.Build(query, args...)
	}
	return w.db.QueryRowContext(ctx, query, args...)
}
