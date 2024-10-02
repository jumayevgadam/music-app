package service

import (
	"context"
	"github.com/jumayevgadam/music-app/internal/database"
	songModel "github.com/jumayevgadam/music-app/internal/models"
	"github.com/jumayevgadam/music-app/pkg/errlst"
	"go.opentelemetry.io/otel"
)

// SongService struct is
type SongService struct {
	repo database.DataStore
}

// NewSongService method is
func NewSongService(repo database.DataStore) *SongService {
	return &SongService{repo: repo}
}

// AddSong service is
func (s *SongService) AddSong(ctx context.Context, dtoModel *songModel.DTO) (int, error) {
	tracer := otel.Tracer("[AddSong][Service]")
	ctx, span := tracer.Start(ctx, "AddSong")
	defer span.End()

	var (
		songID int
		err    error
	)

	if err := s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		songID, err = db.SongRepo().AddSong(ctx, dtoModel.ToStorage())
		if err != nil {
			return errlst.ParseErrors(err)
		}

		return nil
	}); err != nil {
		return -1, errlst.ParseErrors(err)
	}

	return songID, nil
}
