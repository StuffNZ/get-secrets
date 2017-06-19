package s3

import (
	log "github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	//"github.com/aws/aws-sdk-go/service/s3"
)

/*
For each given S3 file, read the contents (into a buffer?)
Pass the buffer to a user-provided function
~Convert the buffer into a gotenv object~
*/

const (
	// BufferSize is the size of Buffer for AWS Download manager to write into
	BufferSize = 1024 * 1024 // 1MB Buffer
)

func (s *Details) Read(subPath string) (string, error) {
	fqPath := s.source.JoinPath(subPath)
	buf, err := s.readWithFqPath(fqPath)
	if err != nil {
		log.WithFields(log.Fields{"full-path": fqPath}).Error(err)
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
