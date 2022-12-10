package util

import (
	"crypto/sha1"
	"fmt"
)

// Hash hashes a string using sha1
func Hash(s string) string {
	hash := sha1.New()
	_, _ = hash.Write([]byte(s))
	bs := hash.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
