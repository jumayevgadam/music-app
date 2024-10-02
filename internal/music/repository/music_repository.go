package repository

import (
	"context"
	"github.com/jumayevgadam/music-app/internal/connection"
	songModel "github.com/jumayevgadam/music-app/internal/models"
	"github.com/jumayevgadam/music-app/pkg/errlst"
)

// SongRepository struct is
type SongRepository struct {
	psqlDB connection.DB
}

// NewSongRepository method is
func NewSongRepository(psqlDB connection.DB) *SongRepository {
	return &SongRepository{psqlDB: psqlDB}
}

// AddSong repo is
func (sr *SongRepository) AddSong(ctx context.Context, daoModel *songModel.DAO) (int, error) {
	var songID int

	if err := sr.psqlDB.QueryRow(
		ctx,
		addSongQuery,
		daoModel.Group,
		daoModel.Title,
		daoModel.ReleaseDate,
		daoModel.Text,
		daoModel.Link,
	).Scan(&songID); err != nil {
		return -1, errlst.ParseSqlErrors(err)
	}

	return songID, nil
}
