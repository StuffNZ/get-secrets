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

// Source Dir for Secret files
type Source struct {
	Base   string
	url    *url.URL
	bucket string
	prefix string
	//session *session.Session
	s3session *s3.S3
}

// Init Source struct
func (s *Source) Init() *Source {
	if err := s.initBase(); err != nil {
		log.Panic(err)
	}
	if err := s.initSession(); err != nil {
		log.Panic(err)
	}

	s.bucket = s.url.Host
	s.prefix = strings.TrimPrefix(s.url.Path, "/") // an initial `/` won't work with S3

	return s
}

func (s *Source) initBase() error {
	var err error

	if s.Base == "" {
		// FIXME: I really want to combine both of these!
		log.Error("No URL provided.")
		return fmt.Errorf("No Base URL provided")
	}
	if s.url, err = url.Parse(s.Base); err != nil {
		log.WithFields(log.Fields{"base-url": s.Base}).Error(err)
		return err
	}

	if s.url.Scheme != "s3" {
		log.WithFields(log.Fields{"url": s.url}).Error("Base URL Scheme is not supported.")
		return fmt.Errorf("Base URL Scheme is not supported: %q", s.url.Scheme)
	}

	return nil
}

func (s *Source) initSession() error {
	// TODO: Enable AWS_SDK_LOAD_CONFIG env-var, somehow!
	session, err := session.NewSession()
	if err != nil {
		log.Error(err)
		return err

	}

	// TODO: This should be in a config file (or ~/.aws/config):
	s.s3session = s3.New(session, &aws.Config{Region: aws.String("ap-southeast-2")})
	return nil
}
