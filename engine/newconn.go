package engine

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewConn(ctx context.Context, cfg *Options) (*engine, error) {
	pool, err := pgxpool.New(ctx, cfg.Addr)
	if err != nil {
		return nil, err
	}

	return &engine{
		pool:   pool,
		config: *cfg,
	}, nil
}
