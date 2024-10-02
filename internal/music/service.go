package music

import (
	"context"
	songModel "github.com/jumayevgadam/music-app/internal/models"
)

// write needed methods for service layer

// Service is
type Service interface {
	AddSong(ctx context.Context, dtoModel *songModel.DTO) (int, error)
}
