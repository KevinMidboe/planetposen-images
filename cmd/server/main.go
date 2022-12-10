package main

import (
	"context"

	"github.com/kevinmidboe/planetposen-images/config"
	"github.com/kevinmidboe/planetposen-images/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	// log.SetFormatter(logrustic.NewFormatter("planetposen-images"))

	log.Info("Starting...")

	ctx := context.Background()
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatal(err.Error())
	}

	var s server.Server

	if err := s.Create(ctx, config); err != nil {
		log.Fatal(err.Error())
	}

	if err := s.Serve(ctx); err != nil {
		log.Fatal(err.Error())
	}
}
