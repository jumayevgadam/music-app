package server

import (
	"github.com/jumayevgadam/music-app/internal/config"
	"github.com/jumayevgadam/music-app/internal/database"
	"github.com/jumayevgadam/music-app/pkg/errlst"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// Server struct keeps all needed configurations for project
type Server struct {
	Echo      *echo.Echo
	Cfg       *config.Config
	DataStore database.DataStore
}

// NewServer is
func NewServer(cfg *config.Config, dataStore database.DataStore) *Server {
	server := &Server{
		Echo:      echo.New(),
		Cfg:       cfg,
		DataStore: dataStore,
	}

	return server
}

// Run the application
func (s *Server) Run() error {
	// Call MapHandlers from here
	if err := s.MapHandlers(s.Echo); err != nil {
		logrus.Println("can not map handlers in Run method")
		return errlst.ParseErrors(err)
	}

	// run http port
	return s.Echo.Start(":" + s.Cfg.Server.HttpPort)
}
