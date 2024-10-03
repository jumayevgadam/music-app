package postgres

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jumayevgadam/music-app/internal/connection"
	"github.com/jumayevgadam/music-app/internal/database"
	"github.com/jumayevgadam/music-app/internal/music"
	musicRepository "github.com/jumayevgadam/music-app/internal/music/repository"
	"github.com/jumayevgadam/music-app/pkg/errlst"
	"github.com/sirupsen/logrus"
)

var _ database.DataStore = (*DataStore)(nil)

// DataStore is
type DataStore struct {
	db        connection.DB
	music     music.Repository
	musicInit sync.Once
}

// NewDataStore is
func NewDataStore(db connection.DB) database.DataStore {
	return &DataStore{db: db}
}

// SongRepo is
func (d *DataStore) SongRepo() music.Repository {
	d.musicInit.Do(func() {
		d.music = musicRepository.NewSongRepository(d.db)
	})

	return d.music
}

// WithTransaction method is
func (d *DataStore) WithTransaction(ctx context.Context, transactionFn database.Transaction) error {
	db, ok := d.db.(connection.DBops)
	if !ok {
		return fmt.Errorf("got error to start transaction")
	}

	// begin transaction
	tx, err := db.Begin(ctx, pgx.TxOptions{})
	if err != nil {
		logrus.Errorf("db.Begin: %v", err)
		return errlst.ParseErrors(err)
	}

	defer func() {
		if err != nil {
			// RollBack the transaction if an error occured
			if err = tx.Rollback(ctx); err != nil {
				logrus.Printf("[postgres][WithTransaction]: failed to rollback transaction: %v", err)
			}
			logrus.Errorf("[postgres][WithTransaction]: transaction failed: %v", err)
		}
	}()

	// transactionalDB is
	transactionalDB := &DataStore{db: tx}
	if err := transactionFn(transactionalDB); err != nil {
		logrus.Println("[postgres][WithTransaction]: transactionFn: %w", err)
		return errlst.ParseErrors(err)
	}

	// Commit the transaction if no error occurred during the transactionFn execution
	if err := tx.Commit(ctx); err != nil {
		logrus.Println("[postgres][WithTransaction]: tx.Commit: %w", err)
		return errlst.ParseErrors(err)
	}

	return nil
}
