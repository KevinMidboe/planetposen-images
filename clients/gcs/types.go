package gcs

import "encoding/base64"

// RawPath is a raw path to a file in GCS
type RawPath string

// EncodedPath is base64-encoded version of RawPath
type EncodedPath string

// Encode base64-encodes raw paths
func (p RawPath) Encode() EncodedPath {
	return EncodedPath(base64.StdEncoding.EncodeToString([]byte(p)))
}

// Decode base64-decodes encoded paths
func (p EncodedPath) Decode() (RawPath, error) {
	decoded, err := base64.StdEncoding.DecodeString(string(p))
	if err != nil {
		return "", err
	}
	return RawPath(decoded), nil
}
