package engine

import "time"

type Options struct {
	Addr             string
	MaxRows          int
	ReadOnly         bool
	StatementTimeout time.Duration
}
