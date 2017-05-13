package s3

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/service/s3"
)

// List all the objects beneath the S3 path
func (s Details) List() ([]string, error) {
	resp, err := s.s3ListObjectsOutput()
	if err != nil {
		return nil, err
	}

	return s.s3MungeListObjectsOutput(s.s3PrefixDir(), resp), nil
}

func (s Details) s3ListObjectsOutput() (*s3.ListObjectsOutput, error) {
	bucket, prefix := (*s.source).Bucket(), (*s.source).Prefix()
	params := &s3.ListObjectsInput{
		Bucket: &bucket,
		Prefix: &prefix,
	}

	resp, err := s.s3Session.ListObjects(params)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	log.WithFields(log.Fields{"s3.params": params, "s3.resp": resp}).Debug()

	return resp, nil
}

func (s Details) s3PrefixDir() string {
	prefixDir := fmt.Sprintf("%s/", (*s.source).Prefix())
	log.WithFields(log.Fields{"s3PrefixDir": prefixDir}).Debug()

	return prefixDir
}

func (s Details) s3MungeListObjectsOutput(prefixDir string, resp *s3.ListObjectsOutput) []string {
	var paths []string

	for _, key := range resp.Contents {
		log.WithFields(log.Fields{"key": *key.Key}).Debug("Found item.")
		if prefixDir != *key.Key {
			paths = append(paths, *key.Key)
		}
	}

	return paths
}
