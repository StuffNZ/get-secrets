package s3

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/hashicorp/go-multierror"
	log "github.com/sirupsen/logrus"
)

/*
For each given S3 file, read the contents (into a buffer?)
Pass the buffer to a user-provided function
~Convert the buffer into a gotenv object~
*/

// BufferSize is the size of Buffer for AWS Download manager to write into
const BufferSize = 1024 * 1024 // 1MB Buffer

// ReadCallback TODO
type ReadCallback func(string, string) error

// ReadList TODO
func (s *Details) ReadList(subPaths []string, f ReadCallback) error {
	var errs *multierror.Error

	for _, subPath := range subPaths {
		// TODO: Should we fail-fast, as soon as an error happens, or aggregate?
		if body, err := s.ReadToString(subPath); err != nil {
			errs = multierror.Append(errs, err)
		} else {
			if errCallback := f(subPath, body); errCallback != nil {
				errs = multierror.Append(errs, errCallback)
			}
		}
	}

	return errs.ErrorOrNil()
}

// ReadToString reads the object at 'subPath' within the
func (s *Details) ReadToString(subPath string) (string, error) {
	if subPath == "" {
		return "", fmt.Errorf("Path %#v is not valid", subPath)
	}
	fqPath := s.source.JoinPath(subPath)
	buf, err := s.readWithFqPath(fqPath)
	if err != nil {
		return "", err
	}
	bufBytes := buf.Bytes()
	bufString := string(bufBytes[:])

	return bufString, nil
}

func (s *Details) readWithFqPath(path string) (*aws.WriteAtBuffer, error) {
	bucket := s.source.Bucket()
	buf := make([]byte, BufferSize)
	writeBuf := aws.NewWriteAtBuffer(buf)

	downloader := s3manager.NewDownloaderWithClient(s.S3())

	params := &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &path,
	}
	_, err := downloader.Download(writeBuf, params)
	if err != nil {
		log.WithFields(log.Fields{"s3.params": params}).Error(err)
		return nil, err
	}

	return writeBuf, nil
}
