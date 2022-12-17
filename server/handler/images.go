package handler

import (
	"encoding/json"
	"github.com/kevinmidboe/planetposen-images/util"
	"strings"

	// "github.com/sirupsen/logrus"
	// "encoding/json"
	"fmt"
	"github.com/kevinmidboe/planetposen-images/clients/gcs"
	"github.com/kevinmidboe/planetposen-images/image"
	// "github.com/dbmedialab/dearheart/event"
	// "github.com/dbmedialab/dearheart/server/internal/serverutils"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"path/filepath"
	// "strconv"
	// "strings"
)

// UploadImages takes a request with file form and uploads the content to GCS
func UploadImages(hostname string, gcsClient gcs.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get initial protocol data
		ctx := r.Context()

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			handleError(w, err, "unable to handle file", http.StatusBadRequest, true)
			return
		}
		var maxSize int64 = 10 * 1024 * 1024 * 1024
		if fileHeader.Size > maxSize {
			handleError(w, nil, "File sized %d, larger than the max of %d\", fileHeader.Size, maxSize", http.StatusBadRequest, true)
			return
		}

		filename := strings.ReplaceAll(fileHeader.Filename, "/", "-")
		defer file.Close()

		logger.InfoWithFilename("uploading image with filename", filename)
		writer, path, err := gcsClient.FileWriter(ctx, filename)
		if err != nil {
			handleGoogleApiError(w, err, "File unable to write file to gcs", http.StatusServiceUnavailable, true)
			return
		}
		defer writer.Close()

		_, err = io.Copy(writer, file)
		if err != nil {
			handleGoogleApiError(w, err, "Error copying file to GCS", http.StatusInternalServerError, true)
		}

		finalURL := util.ImageURL(hostname, string(path))
		responseStruct := image.Image{
			Path: string(path),
			URL:  finalURL,
		}
		logger.UploadSuccessMessage(string(path), finalURL)

		responseData, _ := json.Marshal(responseStruct)
		_, _ = w.Write(responseData)
	}
}

// FetchImage gets a single image from GCS
func FetchImage(gcsClient gcs.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		path := gcs.EncodedPath(mux.Vars(r)["path"])

		logger.InfoWithPath("Getting image", string(path))

		if path == "" {
			handleError(w, nil, "missing image path ", http.StatusBadRequest, true)
			return
		}

		reader, err := gcsClient.FileReader(ctx, path)
		if err != nil {
			handleGoogleApiError(w, err, "error from gcs file reader ", http.StatusBadRequest, true)
			return
		}
		defer reader.Close()

		// we can ignore the error, because we've already verified the path decodes
		filename, _ := path.Decode()
		extension := filepath.Ext(string(filename))
		if extension != "" {
			w.Header().Set("Content-Type", fmt.Sprintf("image/%s", extension[1:]))
		}

		logger.InfoWithFilename("found and returning file from bucket", string(filename))

		_, err = io.Copy(w, reader)
		if err != nil {
			handleError(w, err, "Couldn't copy the file from GCS ", http.StatusInternalServerError, true)
		}
	}
}

func ListImages(gcsClient gcs.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Listing images")
		ctx := r.Context()

		files, err := gcsClient.FileLister(ctx)
		if err != nil {
			handleGoogleApiError(w, err, "error from gcs file lister ", http.StatusBadRequest, true)
			return
		}

		w.WriteHeader(http.StatusOK)
		responseJSON, _ := json.Marshal(struct {
			Message string   `json:"message"`
			Success bool     `json:"success"`
			Files   []string `json:"files"`
		}{
			Message: "Google storage bucket contents",
			Success: true,
			Files:   files,
		})
		w.Write(responseJSON)
	}
}
