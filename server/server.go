// Package server provides functionality to easily set up an HTTTP server.
//
// Clients:
//
//	Database
package server

import (
	"context"
	"fmt"
	log "github.com/kevinmidboe/planetposen-images/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/kevinmidboe/planetposen-images/clients/gcs"
	"github.com/kevinmidboe/planetposen-images/config"
)

// Server holds the HTTP server, router, config and all clients.
type Server struct {
	Config    *config.Config
	HTTP      *http.Server
	Router    *mux.Router
	GCSClient gcs.Client
}

var logger = log.InitLogger()

// Create sets up the HTTP server, router and all clients.
// Returns an error if an error occurs.
func (s *Server) Create(ctx context.Context, config *config.Config) error {
	// metrics.RegisterPrometheusCollectors()

	s.Config = config
	s.Router = mux.NewRouter()
	s.HTTP = &http.Server{
		Addr:    fmt.Sprintf(":%s", s.Config.Port),
		Handler: s.Router,
	}

	gcsClient, err := gcs.NewClient(ctx, config.GCSBucket)
	if err != nil {
		return err
	}

	s.GCSClient = gcsClient

	s.setupRoutes()

	return nil
}

// Serve tells the server to start listening and serve HTTP requests.
// It also makes sure that the server gracefully shuts down on exit.
// Returns an error if an error occurs.
func (s *Server) Serve(ctx context.Context) error {
	// closer, err := trace.InitGlobalTracer(s.Config)

	// if err != nil {
	// 	return err
	// }

	// defer closer.Close()

	go func(ctx context.Context, s *Server) {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

		<-stop

		logger.Info("Shutdown signal received")

		if err := s.HTTP.Shutdown(ctx); err != nil {
			logger.Error("Error causing shutdown", err)
		}
	}(ctx, s)

	logger.Info("Ready at: " + s.Config.Port)

	if err := s.HTTP.ListenAndServe(); err != http.ErrServerClosed {
		logger.Fatal(err)
	}

	return nil
}
