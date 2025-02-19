package storage

import (
	"context"
	"database/sql"
)

type storage struct {
	db DBTX
}

type DBTX interface {
	ExecContext(context.Context, string, ...any) (sql.Result, error)
	QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...any) *sql.Row
}

func New(db DBTX) *storage {
	return &storage{db}
}
