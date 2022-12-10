package gcs

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/kevinmidboe/planetposen-images/util"
	"google.golang.org/api/iterator"
	"path/filepath"
	// "errors"
	"fmt"
	// "github.com/dbmedialab/dearheart/event"
	// "github.com/dbmedialab/dearheart/util"
	// "path/filepath"
	// "time"
)

// Client represents a GCS client with the functions that *we* need.
type Client interface {
	FileWriter(ctx context.Context, filename string) (*storage.Writer, EncodedPath, error)
	FileReader(ctx context.Context, path EncodedPath) (*storage.Reader, error)
	FileLister(ctx context.Context) ([]string, error)
}

// NewClient instantiates our default GCS client.
func NewClient(ctx context.Context, bucketName string) (Client, error) {
	c, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	return &clientImpl{
		BucketName: bucketName,
		GCSClient:  c,
	}, nil
}

type clientImpl struct {
	BucketName string
	GCSClient  *storage.Client
}

func (c *clientImpl) FileWriter(ctx context.Context, filename string) (writer *storage.Writer, path EncodedPath, err error) {
	extension := filepath.Ext(filename)
	rawPath := RawPath(util.Hash(filename) + extension)
	bucket := c.GCSClient.Bucket(c.BucketName)
	object := bucket.Object(string(rawPath))

	path = rawPath.Encode()
	writer = object.NewWriter(ctx)
	return
}

func (c *clientImpl) FileReader(ctx context.Context, path EncodedPath) (reader *storage.Reader, err error) {
	decoded, err := path.Decode()
	if err != nil {
		return nil, fmt.Errorf("invalid path %s", err)
	}

	bucket := c.GCSClient.Bucket(c.BucketName)
	object := bucket.Object(string(decoded))

	reader, err = object.NewReader(ctx)
	return
}

func (c *clientImpl) FileLister(ctx context.Context) (files []string, err error) {
	query := &storage.Query{Prefix: ""}
	var names []string
	bucket := c.GCSClient.Bucket(c.BucketName)
	it := bucket.Objects(ctx, query)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error from query %s", err)
		}
		names = append(names, attrs.Name)
	}

	return names, nil
}
