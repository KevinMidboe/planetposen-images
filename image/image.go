package image

import (
	"github.com/kevinmidboe/planetposen-images/util"
)

// MessageImage is a representation of a single image in the database
type Image struct {
	Path      string `json:"path"`
	URL       string `json:"url,omitempty"`
	RemoteURL string `json:"remote_url"`
}

// GetURL gets URL of the image, also in cases where MessageImage.URL is not defined.
func (mi *Image) GetURL(hostname string) string {
	if mi.URL != "" {
		return mi.URL
	}
	return util.ImageURL(hostname, mi.Path)
}

type PostImageData struct {
	// PageTitle string
	Filename string
}
