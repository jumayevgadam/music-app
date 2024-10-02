package routes

import (
	"github.com/jumayevgadam/music-app/internal/database"
	"github.com/jumayevgadam/music-app/internal/music/handler"
	"github.com/jumayevgadam/music-app/internal/music/service"
	"github.com/labstack/echo/v4"
)

// We use in routes package needed http routes for songs

// Routes is
func Routes(e *echo.Group, dataStore database.DataStore) {
	// init Service
	Service := service.NewSongService(dataStore)
	// init Handler
	Handler := handler.NewSongHandler(Service)

	// init main group for songs
	songGroup := e.Group("/song")

	// Endpoints are
	{
		songGroup.POST("/create", Handler.AddSong())
	}
}
