package connection

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jumayevgadam/music-app/internal/config"
	"github.com/sirupsen/logrus"
)

// Now we'll write decorators for performing DB operations

// DBops is
var _ DB = (*Database)(nil)

// Querier interface is
type Querier interface {
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
}

var (
	_ Querier = &pgxpool.Pool{}
	_ Querier = &pgxpool.Conn{}
)

// DB interface for general database operations
type DB interface {
	Get(ctx context.Context, db Querier, dest interface{}, query string, args ...interface{}) error
	Select(ctx context.Context, db Querier, dest interface{}, query string, args ...interface{}) error
	QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
}

// DBops interface for general database operations with Transaction
type DBops interface {
	DB
	Begin(ctx context.Context, txOpts pgx.TxOptions) (TxOps, error)
	Close() error
}

// Database struct implementing the DBops interface
type Database struct {
	db *pgxpool.Pool
}

// GetDBClient initializes and returns a new Database instance
func GetDBClient(ctx context.Context, cfg config.Postgres) (*Database, error) {
	db, err := pgxpool.New(ctx, fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.SslMode,
	))

	if err != nil {
		return nil, fmt.Errorf("connection.GetDBClient: %w", err)
	}

	if err := db.Ping(ctx); err != nil {
		return nil, fmt.Errorf("connection.GetDBClient.Ping: %w", err)
	}

	return &Database{db: db}, nil
}

// Get implements the DB interface
func (d *Database) Get(ctx context.Context, db Querier, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Get(ctx, db, dest, query, args...)
}

// Select retrieves multiple records and scans them into dest
func (d *Database) Select(ctx context.Context, db Querier, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Select(ctx, db, dest, query, args...)
}

// QueryRow executes a query expected to return at most one row
func (d *Database) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return d.db.QueryRow(ctx, query, args...)
}

// Query executes a query that returns multiple rows
func (d *Database) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return d.db.Query(ctx, query, args...)
}

// Exec executes a query that doesn't return rows
func (d *Database) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return d.db.Exec(ctx, query, args...)
}

// Begin starts a new transaction
func (d *Database) Begin(ctx context.Context, txOpts pgx.TxOptions) (TxOps, error) {
	if d == nil {
		return nil, fmt.Errorf("cannot start transaction")
	}

	c, err := d.db.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	tx, err := d.db.BeginTx(ctx, txOpts)
	if err != nil {
		c.Release()
		logrus.Printf("Failed to begin transaction: %v", err)
		return nil, fmt.Errorf("connection.Database.Begin: %w", err)
	}
	return &Transaction{Tx: tx}, nil
}

// Close closes the database connection pool
func (d *Database) Close() error {
	d.db.Close()
	return nil
}
