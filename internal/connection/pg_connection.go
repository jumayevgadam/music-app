package connection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jumayevgadam/music-app/internal/config"
)

// Now we'll write decorators for performing DB operations

// DBops is
var _ DB = (*Database)(nil)

// Querier is
type Querier interface {
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
}

// Querier
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
func GetDBClient(ctx context.Context, cfg *config.Postgres) (*Database, error) {
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
