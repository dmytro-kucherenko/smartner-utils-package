package queries

import (
	"context"
	"database/sql"
)

type manager[T Prepared] struct {
	connection  *sql.DB
	queries     T
	transaction *sql.Tx
	new         func(*sql.DB) T
	prepare     func(context.Context, *sql.DB) (T, error)
}

func NewSQLManager[T Prepared](
	connection *sql.DB,
	new func(*sql.DB) T,
	prepare func(context.Context, *sql.DB) (T, error),
) Manager[T] {
	return &manager[T]{connection: connection, new: new, prepare: prepare, queries: new(connection)}
}

func (service *manager[T]) Queries() T {
	return service.queries
}

func (service *manager[T]) Prepare(ctx context.Context) error {
	prepared, err := service.prepare(ctx, service.connection)
	if err != nil {
		return err
	}

	service.queries = prepared

	return nil
}

func (manager *manager[T]) Transaction() *sql.Tx {
	return manager.transaction
}

func (manager *manager[T]) SetTransaction(transaction *sql.Tx) {
	manager.transaction = transaction
}

func (manager *manager[T]) Begin() (*sql.Tx, error) {
	transaction, err := manager.connection.Begin()
	if err != nil {
		return nil, err
	}

	manager.SetTransaction(transaction)

	return transaction, nil
}

func (manager *manager[T]) Close() error {
	manager.transaction = nil

	err := manager.queries.Close()
	manager.queries = manager.new(manager.connection)

	return err
}
