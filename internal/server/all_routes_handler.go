package server

import (
	songHttp "github.com/jumayevgadam/music-app/internal/music/routes"
	"github.com/labstack/echo/v4"
)

const (
	v1URL = "/api/v1"
)

// MapHandlers is
func (s *Server) MapHandlers(e *echo.Echo) error {
	//* v1 is
	v1 := s.Echo.Group(v1URL)
	// song-http route is
	songHttp.Routes(v1, s.DataStore)

	return nil
}
