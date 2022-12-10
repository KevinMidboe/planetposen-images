package util

import "fmt"

// ImageURL creates imageURL from hostname and image name
func ImageURL(hostname, name string) string {
	return fmt.Sprintf("https://%s/api/v1/images/%s", hostname, name)
}
