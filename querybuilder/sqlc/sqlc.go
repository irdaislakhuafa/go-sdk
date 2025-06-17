package sqlc

import (
	"context"
	"database/sql"
	"strings"

	"github.com/irdaislakhuafa/go-sdk/collections"
	"github.com/irdaislakhuafa/go-sdk/convert"
)

type (
	DBTX interface {
		ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
		PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
		QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
		QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	}

	ctxKey struct{}

	BuilderInterface interface {
		Where(expression string, args ...any) BuilderInterface
		In(column string, args ...any) BuilderInterface
		Or(column string, args ...any) BuilderInterface
		And(column string, args ...any) BuilderInterface
		Order(cols string, args ...any) BuilderInterface
		Limit(limit int) BuilderInterface
		Offset(offset int) BuilderInterface
		GroupBy(cols string, args ...any) BuilderInterface
		Build(query string, args ...any) (string, []any)
	}

	Builder struct {
		filters       []filter
		order         order
		offset, limit *int
		groupBy       *groupBy
	}

	order struct {
		expression string
		args       []any
	}

	filter struct {
		expression string
		hasLogic   bool
		args       []any
	}

	groupBy struct {
		cols string
		args []any
	}
)

// Build create new builder or clone builder from context
func Build(ctx context.Context, fn func(b *Builder)) context.Context {
	b, isOk := GetBuilder(ctx)
	if !isOk {
		b = &Builder{}
	} else {
		b = convert.ToPointer(*b)
	}

	fn(b)
	return WithBuilder(ctx, b)
}

func GetBuilder(ctx context.Context) (*Builder, bool) {
	b, isOk := ctx.Value(ctxKey{}).(*Builder)
	return b, isOk
}

func WithBuilder(ctx context.Context, b *Builder) context.Context {
	return context.WithValue(ctx, ctxKey{}, b)
}

func GenQueryArgs[T any](ctx context.Context, params ...T) (exprs string, args []any) {
	exprs = strings.Repeat("?,", len(params))
	exprs = exprs[:len(exprs)-1]
	args = collections.Map(params, func(i int, v T) any { return v })
	return exprs, args
}
