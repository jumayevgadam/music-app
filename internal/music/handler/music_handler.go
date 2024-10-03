package handler

import (
	"net/http"

	songModel "github.com/jumayevgadam/music-app/internal/models"
	musicOps "github.com/jumayevgadam/music-app/internal/music"
	httpError "github.com/jumayevgadam/music-app/pkg/errlst"
	"github.com/jumayevgadam/music-app/pkg/reqvalidator"
	"github.com/jumayevgadam/music-app/pkg/tracing"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
)

// SongHandler struct is
type SongHandler struct {
	service musicOps.Service
}

// NewSongHandler method is
func NewSongHandler(service musicOps.Service) *SongHandler {
	return &SongHandler{service: service}
}

func (sh *SongHandler) AddSong() echo.HandlerFunc {
	return func(c echo.Context) error {
		tracer := otel.Tracer("[SongHandler][AddSong]")
		ctx, span := tracer.Start(c.Request().Context(), "[SongHandler][AddSong]")
		defer span.End()

		var songRequest songModel.DTO
		if err := reqvalidator.ReadRequest(c, &songRequest); err != nil {
			tracing.EventErrorTracer(span, err, "[SongHandler][AddSong]")
			return c.JSON(httpError.Response(err))
		}

		songID, err := sh.service.AddSong(ctx, &songRequest)
		if err != nil {
			tracing.EventErrorTracer(span, err, "[SongHandler][AddSong]")
			return c.JSON(httpError.Response(err))
		}

		return c.JSON(http.StatusOK, songID)
	}
}
