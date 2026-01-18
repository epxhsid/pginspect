package engine

import (
	"context"
	"fmt"
	"time"
)

type TableData struct {
	Columns  []string
	Rows     [][]any
	RowCount int
}

func (e *engine) TableData(ctx context.Context, schema, table string, limit int) (*TableData, error) {
	if limit <= 0 || limit > e.config.MaxRows {
		limit = e.config.MaxRows
	}

	timeout := e.config.StatementTimeout
	if timeout <= 0 {
		timeout = 2 * time.Second
	}

	query := fmt.Sprintf(`SELECT * FROM "%s"."%s" LIMIT %d`, schema, table, limit)

	rows, err := e.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query table data: %w", err)
	}
	defer rows.Close()

	colNames := rows.FieldDescriptions()
	columns := make([]string, len(colNames))
	for i, col := range colNames {
		columns[i] = string(col.Name)
	}

	var data [][]any
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		data = append(data, values)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return &TableData{
		Columns:  columns,
		Rows:     data,
		RowCount: len(data),
	}, nil
}
