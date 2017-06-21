package main

import (
	s3ish "build-dotenv/files/s3"
	urlish "build-dotenv/files/s3/s3url"
	"fmt"
	log "github.com/sirupsen/logrus"
	// "build-dotenv/cmd"
)

func main() {
	log.Debug("Starting...")

	s3url := urlish.New().WithURL("s3://kiwiops-ecs-staging-env/stuff-brightcove-video-service")
	s3 := s3ish.New().WithSource(s3url)
	s3lists, err := s3.List()
	if err != nil {
		log.Panic(err)
	}
	s3lists = append(s3lists, "")
	s3lists = append(s3lists, "")
	// s3lists, _ := (&s3ish.Source{URLString: "s3://kiwiops-ecs-staging-env/stuff-brightcove-video-service"}).Init().List()
	log.WithFields(log.Fields{"list": s3lists}).Info()

	// s3lists, _ := (&s3ish.Source{URLString: "s3://kiwiops-ecs-staging-env/stuff-brightcove-video-service"}).Init().List()
	if errs := s3.ReadList(s3lists, readCallback); errs != nil {
		log.Panic(errs)
	}
}

func readCallback(path, body string) error {
	if path == "" {
		log.Error("Empty path!")
		return fmt.Errorf("Empty path")
	}
	if body == "" {
		log.Error("Empty body!")
		return fmt.Errorf("Empty body")
	}
	log.WithFields(log.Fields{"path": path}).Info(body)
	return nil
}
