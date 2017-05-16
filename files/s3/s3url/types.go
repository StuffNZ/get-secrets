package s3url

import (
	"fmt"
	"net/url"
	"strings"

	log "github.com/Sirupsen/logrus"
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
	clone := *s // This does a shallow clone

	if url == "" {
		log.Panic("No Base URL provided")
	}

	clone.urlString = url
	clone.url = clone.urlFromURLString(clone.urlString)
	clone.bucket, clone.prefix = s.bucketPrefixFromURL(clone.url)

	return &clone
}

// WithBucket creates new struct with `Bucket` updated and `Prefix` empty
func (s *Path) WithBucket(bucket string) *Path {
	return s.WithBucketPrefix(bucket, "")
}

// WithBucketPrefix creates new struct with `Bucket` and `Prefix` updated
func (s *Path) WithBucketPrefix(bucket, prefix string) *Path {
	clone := *s // This does a shallow clone

	if bucket == "" {
		log.Panic("Empty Bucket provided")
	}

	if prefix == "" {
		prefix = "/"
	}

	clone.bucket, clone.prefix = bucket, prefix
	clone.urlString = fmt.Sprintf("s3://%s/%s", bucket, prefix)
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

	bucket := url.Host
	// an initial `/` won't work with S3
	prefix := strings.TrimPrefix(url.Path, "/")

	return bucket, prefix
}
