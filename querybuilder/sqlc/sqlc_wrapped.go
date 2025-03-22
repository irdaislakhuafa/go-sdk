package sqlc

import (
	"context"
	"database/sql"
	"fmt"
)

type (
	wrappedDBTX struct {
		db   DBTX
		opts WrappedOpts
	}

	WrappedOpts struct {
		showQuery bool
		showArgs  bool
	}
)

func Wrap(db DBTX, opts WrappedOpts) DBTX {
	return &wrappedDBTX{
		db:   db,
		opts: opts,
	}
}

func (w *wrappedDBTX) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	if b, isOk := GetBuilder(ctx); isOk {
		query, args = b.Build(query, args...)
	}

	if w.opts.showQuery {
		fmt.Printf("query: %v\n", query)
	}

	if w.opts.showArgs {
		fmt.Printf("args: %v\n", args)
	}

	return w.db.ExecContext(ctx, query, args...)
}

func (w *wrappedDBTX) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	if w.opts.showQuery {
		fmt.Printf("query: %v\n", query)
	}

	return w.db.PrepareContext(ctx, query)
}

func (w *wrappedDBTX) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	if b, isOk := GetBuilder(ctx); isOk {
		query, args = b.Build(query, args...)
	}

	if w.opts.showQuery {
		fmt.Printf("query: %v\n", query)
	}

	if w.opts.showArgs {
		fmt.Printf("args: %v\n", args)
	}

	return w.db.QueryContext(ctx, query, args...)
}

func (w *wrappedDBTX) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	if b, isOk := GetBuilder(ctx); isOk {
		query, args = b.Build(query, args...)
	}

	if w.opts.showQuery {
		fmt.Printf("query: %v\n", query)
	}

	if w.opts.showArgs {
		fmt.Printf("args: %v\n", args)
	}

	return w.db.QueryRowContext(ctx, query, args...)
}
