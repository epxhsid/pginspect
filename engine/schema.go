package engine

import (
	"context"
	"fmt"
	"time"
)

func (e *engine) Schemas(ctx context.Context) ([]string, error) {
	timeout := e.config.StatementTimeout
	if timeout <= 0 {
		timeout = 3 * time.Second
	}

	query := `
		SELECT schema_name
		FROM information_schema.schemata
		ORDER BY schema_name ASC
	`

	rows, err := e.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query schemas: %w", err)
	}

	defer rows.Close()
	var schemas []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, fmt.Errorf("failed to scan schema name: %w", err)
		}
		schemas = append(schemas, name)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over schema rows: %w", err)
	}

	return schemas, nil

}
