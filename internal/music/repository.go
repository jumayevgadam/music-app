package music

import (
	"context"
	songModel "github.com/jumayevgadam/music-app/internal/models"
)

// write needed methods for repository layer

// Repository is
type Repository interface {
	AddSong(ctx context.Context, daoModel *songModel.DAO) (int, error)
	//GetSongByID(ctx context.Context, id int) (*songModel.DAO, error)
	//GetSongByName(ctx context.Context, name string) (*songModel.DAO, error)
	//GetAllSongs(ctx context.Context) ([]*songModel.DAO, error)
	//UpdateSong(ctx context.Context, songID int) (string, error)
	//DeleteSong(ctx context.Context, songID int) (string, error)
}
