package url

import (
	"net/url"
	"strings"

	log "github.com/Sirupsen/logrus"
)

// Source Dir for Secret files
type Config struct {
	URL    string
	Bucket string
	Prefix string
}

type self struct {
	config    Config
	urlString string
	url       *url.URL
	bucket    string
	prefix    string
}

func New(c Config) *self {
	s := &self{config: c}

	if c.Bucket == "" {
		if c.Prefix == "" {
			if c.URL == "" {
				log.Panic("No Base URL provided.")
			}
			s.bucket, s.prefix = s.bucketPrefixFromURL(s.urlFromURLString(c.URL))

		} else {
			log.Panic("Empty Bucket provided with Prefix.")
		}

	} else {
		s.bucket = c.Bucket
		s.prefix = c.Prefix

		if s.prefix == "" {
			s.prefix = "/"
		}
	}

	return s
}

func (s *self) Bucket() string {
	return s.bucket
}

func (s *self) Prefix() string {
	return s.prefix
}

func (s *self) urlFromURLString(urlString string) *url.URL {
	url, err := url.Parse(urlString)
	if err != nil {
		log.WithFields(log.Fields{"url-string": urlString}).Panic(err)
	}

	return url
}

func (s *self) bucketPrefixFromURL(url *url.URL) (string, string) {
	if url.Scheme != "s3" {
		log.WithFields(log.Fields{"url": url}).Panic("Base URL Scheme is not supported.")
	}

	bucket := url.Host
	// an initial `/` won't work with S3
	prefix := strings.TrimPrefix(url.Path, "/")

	return bucket, prefix
}
