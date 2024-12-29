package queries

import (
	"context"
	"database/sql"
)

type DB[T Prepared] interface {
	New(*sql.DB) T
	Prepare(context.Context, *sql.DB) (T, error)
}

type Prepared interface {
	Close() error
}

type TransactionCaller[T Prepared] interface {
	Begin() (*sql.Tx, error)
	SetTransaction(*sql.Tx)
	Close() error
}

type PrepareCaller[T Prepared] interface {
	Prepare(context.Context) error
	Close() error
}

type Manager[T Prepared] interface {
	TransactionCaller[T]
	PrepareCaller[T]
	Queries() T
	Transaction() *sql.Tx
}
