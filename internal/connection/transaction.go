package connection

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// TxOps interface for transaction operations
type TxOps interface {
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	DB
}

// Transaction struct for handling transactions
type Transaction struct {
	Tx   pgx.Tx
	Conn *pgxpool.Conn
}

// Get retrieves a single record and scans it into dest
func (tx *Transaction) Get(ctx context.Context, db Querier, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Get(ctx, tx.Tx, dest, query, args...)
}

// Select retrieves multiple records and scans them into dest
func (tx *Transaction) Select(ctx context.Context, db Querier, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Select(ctx, tx.Tx, dest, query, args...)
}

// QueryRow executes a query expected to return at most one row
func (tx *Transaction) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return tx.Tx.QueryRow(ctx, query, args...)
}

// Query executes a query that returns multiple rows
func (tx *Transaction) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return tx.Tx.Query(ctx, query, args...)
}

// Exec executes a query that doesn't return rows
func (tx *Transaction) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return tx.Tx.Exec(ctx, query, args...)
}

// Commit commits the transaction
func (tx *Transaction) Commit(ctx context.Context) error {
	return tx.Tx.Commit(ctx)
}

// Rollback rolls back the transaction
func (tx *Transaction) Rollback(ctx context.Context) error {
	return tx.Tx.Rollback(ctx)
}
