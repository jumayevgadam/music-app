package music

import "github.com/labstack/echo/v4"

// write needed methods for Handler layer

// Handler interface is
type Handler interface {
	AddSong() echo.HandlerFunc
}
