package main

import (
	"context"
	"github.com/kevinmidboe/planetposen-images/config"
	log "github.com/kevinmidboe/planetposen-images/logger"
	"github.com/kevinmidboe/planetposen-images/server"
)

func main() {
	logger := log.InitLogger()
	logger.Info("Starting...")

	ctx := context.Background()
	config, err := config.LoadConfig()

	if err != nil {
		logger.Fatal(err)
	}

	var s server.Server

	if err := s.Create(ctx, config); err != nil {
		logger.Fatal(err)
	}

	if err := s.Serve(ctx); err != nil {
		logger.Fatal(err)
	}
}
