package handler

import (
	"encoding/json"
	log "github.com/kevinmidboe/planetposen-images/logger"
	"net/http"
)

var logger = log.InitLogger()

// handleError - Logs the error (if shouldLog is true), and outputs the error message (msg)
func handleError(w http.ResponseWriter, err error, msg string, statusCode int, shouldLog bool) {
	if shouldLog {
		logger.Error(msg, err)
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

func handleGoogleApiError(w http.ResponseWriter, err error, msg string, statusCode int, shouldLog bool) {
	if shouldLog {
		logger.GoogleApiError(msg, err)
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
