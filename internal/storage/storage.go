package storage

import (
	"context"
	"database/sql"
)

type (
	DBTX interface {
		ExecContext(context.Context, string, ...any) (sql.Result, error)
		QueryContext(context.Context, string, ...any) (*sql.Rows, error)
		QueryRowContext(context.Context, string, ...any) *sql.Row
	}

	storage struct {
		db DBTX
	}
)

func New(db DBTX) *storage {
	return &storage{db}
}
