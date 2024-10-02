package main

import (
	"context"
	"github.com/jumayevgadam/music-app/internal/config"
	"github.com/jumayevgadam/music-app/internal/connection"
	"github.com/jumayevgadam/music-app/internal/database/postgres"
	"github.com/jumayevgadam/music-app/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		logrus.Errorf("[main][LoadConfig]: %v", err.Error())
	}

	// implement tracer here

	// implement psqlDB
	psqlDB, err := connection.GetDBClient(context.Background(), cfg.Postgres)
	if err != nil {
		logrus.Errorf("[main][GetDBClient]: %v", err.Error())
	} else {
		logrus.Println("Connected to PostgreSQL")
	}

	defer func() {
		if err := psqlDB.Close(); err != nil {
			logrus.Errorf("[main][Close]: %v", err.Error())
		}
	}()

	// implement dataStore here
	dataStore := postgres.NewDataStore(psqlDB)

	// source is
	source := server.NewServer(cfg, dataStore)
	if err := source.Run(); err != nil {
		logrus.Errorf("[main][Run]: %v", err.Error())
	}
}
