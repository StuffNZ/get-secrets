package s3

import (
	"fmt"
	// "strings"
	// "net/url"

	log "github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/service/s3"
)

// List all the objects beneath the S3 path
func (s Source) List() ([]string, error) {
	s3PrefixDir := fmt.Sprintf("%s/", s.prefix)
	log.WithFields(log.Fields{"s3PrefixDir": s3PrefixDir}).Debug()

	params := &s3.ListObjectsInput{
		Bucket: &s.bucket,
		Prefix: &s.prefix,
	}
	resp, err := s.s3session.ListObjects(params)
	if err != nil {
		log.Fatal(err)
		return nil, err

	}
	var paths []string

	log.WithFields(log.Fields{"s3.params": params, "s3.resp": resp}).Debug()

	for _, key := range resp.Contents {
		log.WithFields(log.Fields{"key": *key.Key}).Debug("Found item.")
		if s3PrefixDir != *key.Key {
			paths = append(paths, *key.Key)
		}
	}

	return paths, nil
}
