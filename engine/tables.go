package engine

import (
	"context"
	"fmt"
	"time"
)

func (e *engine) Tables(ctx context.Context, schema string) ([]string, error) {
	timeout := e.config.StatementTimeout
	if timeout <= 0 {
		timeout = 2 * time.Second
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	query := `
		SELECT table_name
		FROM information_schema.tables
		WHERE table_schema = $1
		ORDER BY table_name
	`

	rows, err := e.pool.Query(ctx, query, schema)
	if err != nil {
		return nil, fmt.Errorf("failed to query tables for schema %q: %w", schema, err)
	}

	defer rows.Close()

	var tables []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return nil, fmt.Errorf("failed to scan table name: %w", err)
		}
		tables = append(tables, tableName)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over table rows: %w", err)
	}

	return tables, nil
}
