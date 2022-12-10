package handler

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// handleError - Logs the error (if shouldLog is true), and outputs the error message (msg)
func handleError(w http.ResponseWriter, err error, msg string, statusCode int, shouldLog bool) {
	if shouldLog {
		log.WithField("err", err).Error(msg)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	errorJSON, _ := json.Marshal(struct {
		Error string `json:"error"`
	}{
		Error: msg,
	})
	w.Write(errorJSON)
}