package engine

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type engine struct {
	pool   *pgxpool.Pool
	config Options
}

type Engine interface {
	Ping(ctx context.Context) error
	Schemas(ctx context.Context) ([]string, error)
	Tables(ctx context.Context, schema string) ([]string, error)
	TableData(ctx context.Context, schema, table string, limit int) (*TableData, error)
}
