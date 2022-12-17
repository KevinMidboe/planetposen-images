package util

import (
	"fmt"
)

// ImageURL creates imageURL from hostname and image name
func ImageURL(hostname, name string) string {
	return fmt.Sprintf("https://%s/api/v1/images/%s", hostname, name)
}

// ImageRemoteURL creates imageURL to bucket file
func ImageRemoteURL(bucketname string, path string) string {
	return fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketname, path)
}
