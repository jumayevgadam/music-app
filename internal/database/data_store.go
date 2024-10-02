package database

import (
	"context"
	"github.com/jumayevgadam/music-app/internal/music"
)

// We want to use clean way implementing 'Transaction' with callback function

// Transaction is
type Transaction func(db DataStore) error

// DataStore is
type DataStore interface {
	WithTransaction(ctx context.Context, tx Transaction) error
	SongRepo() music.Repository
}
