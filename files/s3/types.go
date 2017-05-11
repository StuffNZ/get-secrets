package s3

import (
	"net/url"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Source Dir for Secret files
type Source struct {
	URLString string
	URL       *url.URL
	Bucket    string
	Prefix    string
	//session *session.Session
	S3Session *s3.S3
}

// Init Source struct
func (s *Source) Init() *Source {
	if s.Bucket == "" && s.Prefix == "" {
		if s.URL == nil {
			if s.URLString == "" {
				log.Panic("No Base URL provided.")
			}
			s.InitURLFromURLString()
		}
		s.InitBucketPrefixFromURL()
	}

	if s.S3Session == nil {
		s.InitSessionFromBucketPrefix()
	}

	return s
}

func (s *Source) InitURLFromURLString() *Source {
	var err error
	if s.URL, err = url.Parse(s.URLString); err != nil {
		log.WithFields(log.Fields{"url-string": s.URLString}).Panic(err)
	}

	return s
}

func (s *Source) InitBucketPrefixFromURL() *Source {
	if s.URL.Scheme != "s3" {
		log.WithFields(log.Fields{"url": s.URL}).Panic("Base URL Scheme is not supported.")
	}

	s.Bucket = s.URL.Host
	// an initial `/` won't work with S3
	s.Prefix = strings.TrimPrefix(s.URL.Path, "/")

	return s
}

func (s *Source) InitSessionFromBucketPrefix() *Source {
	// TODO: Enable AWS_SDK_LOAD_CONFIG env-var, somehow!
	session, err := session.NewSession()
	if err != nil {
		log.Panic(err)
	}

	// TODO: `Region` should be in a config file (or ~/.aws/config):
	s.S3Session = s3.New(
		session,
		&aws.Config{Region: aws.String("ap-southeast-2")},
	)
	return s
}
