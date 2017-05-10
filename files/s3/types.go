package s3

import (
	"fmt"
	"net/url"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Source struct {
	Url    string
	url    *url.URL
	bucket string
	prefix string
	//session *session.Session
	s3session *s3.S3
}

func (s *Source) Init() *Source {
	if err := s.initUrl(); err != nil {
		log.Panic(err)
	}
	if err := s.initSession(); err != nil {
		log.Panic(err)
	}

	s.bucket = s.url.Host
	s.prefix = strings.TrimPrefix(s.url.Path, "/") // an initial `/` won't work with S3

	return s
}

func (s *Source) initUrl() error {
	var err error

	if s.Url == "" {
		log.Error("No URL provided.")
		return fmt.Errorf("No URL provided.")
	}
	if s.url, err = url.Parse(s.Url); err != nil {
		log.WithFields(log.Fields{"url": s.Url}).Error(err)
		return err
	}

	if s.url.Scheme != "s3" {
		log.WithFields(log.Fields{"url": s.url}).Error("URL Scheme is not supported.")
		return fmt.Errorf("URL Scheme is not supported: %q", s.url.Scheme)
	}

	return nil
}

func (s *Source) initSession() error {
	// TODO: Enable AWS_SDK_LOAD_CONFIG env-var, somehow!
	if session, err := session.NewSession(); err != nil {
		log.Error(err)
		return err

	} else {
		s.s3session = s3.New(session, &aws.Config{Region: aws.String("ap-southeast-2")})
		return nil
	}
}
