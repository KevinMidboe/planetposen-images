package server

import (
	"github.com/kevinmidboe/planetposen-images/server/handler"
)

const v1API string = "/api/v1"

func (s *Server) setupRoutes() {
	s.Router.HandleFunc("/_healthz", handler.Healthz).Methods("GET").Name("Health")

	api := s.Router.PathPrefix(v1API).Subrouter()
	api.HandleFunc("/images", handler.UploadImages(s.Config.Hostname, s.Config.GCSBucket, s.GCSClient)).Methods("POST").Name("UploadImages")
	api.HandleFunc("/images", handler.ListImages(s.GCSClient)).Methods("GET").Name("ListImages")

	// Raw image fetcher
	api.HandleFunc("/images/{path}", handler.FetchImage(s.GCSClient)).Methods("GET").Name("FetchImage")

}
