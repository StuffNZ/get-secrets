package main

import (
	s3ish "build-dotenv/files/s3"
	urlish "build-dotenv/files/s3/s3url"
	log "github.com/Sirupsen/logrus"
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
	// s3lists, _ := (&s3ish.Source{URLString: "s3://kiwiops-ecs-staging-env/stuff-brightcove-video-service"}).Init().List()
	log.WithFields(log.Fields{"list": s3lists}).Info()
}
