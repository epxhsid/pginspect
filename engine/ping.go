package engine

import (
	"context"
	"fmt"
	"time"
)

func (e *engine) Ping(ctx context.Context) error {
	timeout := e.config.StatementTimeout
	if timeout <= 0 {
		timeout = 2 * time.Second
	}

	// create context with timeout
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// Ping the Postgres pool
	if err := e.pool.Ping(ctx); err != nil {
		return fmt.Errorf("ping failed: %w", err)
	}
	return nil
}
