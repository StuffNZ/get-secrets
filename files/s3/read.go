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

// ReadCallback defines the func-type that can be used as a Callback
type ReadCallback func(string, string) error

/* ReadList reads the provided list of paths, and then sends the contents of these to the provided Callback.
   The Callback function is executed once per path (file object) not as an aggregate.
   Any errors found when reading the object, or returned by the Callback *are* aggregated into a "multierror" object.
   If there are errors, the multierror object is returned, otherwise a "nil" is returned.
*/
func (s *Details) ReadList(subPaths []string, f ReadCallback) error {
	// The errors are aggregated into this multierror object:
	var errs *multierror.Error

	for _, subPath := range subPaths {
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

/*
ReadToString reads the object at 'subPath' within the bucket/prefix, and returns the contents as a string.

If there are any errors when reading, they are returned.
*/
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
