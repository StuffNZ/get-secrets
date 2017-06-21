package s3url

import (
	"fmt"
	"net/url"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Path contains the attrs for Secret files
type Path struct {
	urlString string
	url       *url.URL
	bucket    string
	prefix    string
}

// New object
func New() *Path {
	return &Path{}
}

// func (s *Path) WithConfig(configFunc func(*Path)) *Path {
// 	return configFunc(s)
// }

// WithURL creates new struct updated according to the provided URL
func (s *Path) WithURL(url string) *Path {
	var err error

	clone := *s // This does a shallow clone

	if url == "" {
		log.Panic("No Base URL provided")
	}

	clone.urlString = url
	clone.url = clone.urlFromURLString(clone.urlString)
	bucket, prefix := s.bucketPrefixFromURL(clone.url)
	clone.bucket, clone.prefix, err = s.tidyBucketPrefix(bucket, prefix)
	if err != nil {
		log.Panic(err)
	}

	return &clone
}

// WithBucket creates new struct with `Bucket` updated and `Prefix` empty
func (s *Path) WithBucket(bucket string) *Path {
	return s.WithBucketPrefix(bucket, "")
}

// WithBucketPrefix creates new struct with `Bucket` and `Prefix` updated
func (s *Path) WithBucketPrefix(bucket, prefix string) *Path {
	var err error

	clone := *s // This does a shallow clone

	clone.bucket, clone.prefix, err = s.tidyBucketPrefix(bucket, prefix)
	if err != nil {
		log.Panic(err)
	}

	clone.urlString = fmt.Sprintf("s3://%s/%s", clone.bucket, clone.prefix)
	clone.url = clone.urlFromURLString(clone.urlString)

	return &clone
}

// Bucket returns the S3 bucket
func (s *Path) Bucket() string {
	return s.bucket
}

// Prefix returns the S3 path prefix
func (s *Path) Prefix() string {
	return s.prefix
}

// PrefixDir returns the S3 path prefix in a directory-explicit format
func (s *Path) PrefixDir() string {
	return fmt.Sprintf("%s/", s.prefix)
}

// JoinPath returns the provided path appended to the PrefixDir
func (s *Path) JoinPath(path string) string {
	return strings.Join([]string{s.PrefixDir(), path}, "")
}

func (s *Path) urlFromURLString(urlString string) *url.URL {
	url, err := url.Parse(urlString)
	if err != nil {
		log.WithFields(log.Fields{"url-string": urlString}).Panic(err)
	}

	return url
}

func (s *Path) bucketPrefixFromURL(url *url.URL) (string, string) {
	if url.Scheme != "s3" {
		log.WithFields(log.Fields{"url": url}).Panic("Base URL Scheme is not supported.")
	}

	return url.Host, url.Path
}

func (s *Path) tidyBucketPrefix(bucket, prefix string) (string, string, error) {
	if bucket == "" {
		return "", "", fmt.Errorf("Empty Bucket provided")
	}

	// Remove leading and trailing `/` (an initial `/` won't work with S3):
	prefix = strings.TrimRight(prefix, "/")
	prefix = strings.TrimLeft(prefix, "/")

	// Add/restore a single trailing `/`:
	if prefix == "" {
		prefix = "/"
	}

	log.WithFields(log.Fields{"bucket": bucket, "prefix": prefix}).Debug("Tidied Bucket and Prefix")

	return bucket, prefix, nil
}
